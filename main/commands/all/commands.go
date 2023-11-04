package all

import (
	"github.com/xxhanxx/Xray-core/main/commands/all/api"
	"github.com/xxhanxx/Xray-core/main/commands/all/tls"
	"github.com/xxhanxx/Xray-core/main/commands/base"
)

// go:generate go run github.com/xxhanxx/Xray-core/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		// cmdConvert,
		tls.CmdTLS,
		cmdUUID,
		cmdX25519,
	)
}
