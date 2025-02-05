package vmess

import (
	"crypto/hmac"
	"crypto/sha256"
	"hash/crc64"
	"strings"
	"sync"

	"github.com/xxhanxx/Xray-core/common/dice"
	"github.com/xxhanxx/Xray-core/common/protocol"
	"github.com/xxhanxx/Xray-core/proxy/vmess/aead"
)

// TimedUserValidator is a user Validator based on time.
type TimedUserValidator struct {
	sync.RWMutex
	users []*protocol.MemoryUser

	behaviorSeed  uint64
	behaviorFused bool

	aeadDecoderHolder *aead.AuthIDDecoderHolder
}

// NewTimedUserValidator creates a new TimedUserValidator.
func NewTimedUserValidator() *TimedUserValidator {
	tuv := &TimedUserValidator{
		users:             make([]*protocol.MemoryUser, 0, 16),
		aeadDecoderHolder: aead.NewAuthIDDecoderHolder(),
	}
	return tuv
}

func (v *TimedUserValidator) Add(u *protocol.MemoryUser) error {
	v.Lock()
	defer v.Unlock()

	v.users = append(v.users, u)

	account := u.Account.(*MemoryAccount)
	if !v.behaviorFused {
		hashkdf := hmac.New(sha256.New, []byte("VMESSBSKDF"))
		hashkdf.Write(account.ID.Bytes())
		v.behaviorSeed = crc64.Update(v.behaviorSeed, crc64.MakeTable(crc64.ECMA), hashkdf.Sum(nil))
	}

	var cmdkeyfl [16]byte
	copy(cmdkeyfl[:], account.ID.CmdKey())
	v.aeadDecoderHolder.AddUser(cmdkeyfl, u)

	return nil
}

func (v *TimedUserValidator) GetAEAD(userHash []byte) (*protocol.MemoryUser, bool, error) {
	v.RLock()
	defer v.RUnlock()

	var userHashFL [16]byte
	copy(userHashFL[:], userHash)

	userd, err := v.aeadDecoderHolder.Match(userHashFL)
	if err != nil {
		return nil, false, err
	}
	return userd.(*protocol.MemoryUser), true, err
}

func (v *TimedUserValidator) Remove(email string) bool {
	v.Lock()
	defer v.Unlock()

	email = strings.ToLower(email)
	idx := -1
	for i, u := range v.users {
		if strings.EqualFold(u.Email, email) {
			idx = i
			var cmdkeyfl [16]byte
			copy(cmdkeyfl[:], u.Account.(*MemoryAccount).ID.CmdKey())
			v.aeadDecoderHolder.RemoveUser(cmdkeyfl)
			break
		}
	}
	if idx == -1 {
		return false
	}
	ulen := len(v.users)

	v.users[idx] = v.users[ulen-1]
	v.users[ulen-1] = nil
	v.users = v.users[:ulen-1]

	return true
}

func (v *TimedUserValidator) GetBehaviorSeed() uint64 {
	v.Lock()
	defer v.Unlock()

	v.behaviorFused = true
	if v.behaviorSeed == 0 {
		v.behaviorSeed = dice.RollUint64()
	}
	return v.behaviorSeed
}

var ErrNotFound = newError("Not Found")

var ErrTainted = newError("ErrTainted")
