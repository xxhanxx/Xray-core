package internet_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/xxhanxx/Xray-core/common"
	"github.com/xxhanxx/Xray-core/common/buf"
	"github.com/xxhanxx/Xray-core/testing/servers/tcp"
	. "github.com/xxhanxx/Xray-core/transport/internet"
)

func TestTCPFastOpen(t *testing.T) {
	tcpServer := tcp.Server{
		MsgProcessor: func(b []byte) []byte {
			return b
		},
	}
	dest, err := tcpServer.StartContext(context.Background(), &SocketConfig{Tfo: 256})
	common.Must(err)
	defer tcpServer.Close()

	ctx := context.Background()
	dialer := DefaultSystemDialer{}
	conn, err := dialer.Dial(ctx, nil, dest, &SocketConfig{
		Tfo: 1,
	})
	common.Must(err)
	defer conn.Close()

	_, err = conn.Write([]byte("abcd"))
	common.Must(err)

	b := buf.New()
	common.Must2(b.ReadFrom(conn))
	if r := cmp.Diff(b.Bytes(), []byte("abcd")); r != "" {
		t.Fatal(r)
	}
}
