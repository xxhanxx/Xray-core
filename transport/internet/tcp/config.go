package tcp

import (
	"github.com/xxhanxx/Xray-core/common"
	"github.com/xxhanxx/Xray-core/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
