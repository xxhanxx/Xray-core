package udp

import (
	"github.com/xxhanxx/Xray-core/common"
	"github.com/xxhanxx/Xray-core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
