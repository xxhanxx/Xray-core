package yaml

import (
	"io"

	"github.com/xxhanxx/Xray-core/common"
	"github.com/xxhanxx/Xray-core/common/cmdarg"
	"github.com/xxhanxx/Xray-core/core"
	"github.com/xxhanxx/Xray-core/infra/conf"
	"github.com/xxhanxx/Xray-core/infra/conf/serial"
	"github.com/xxhanxx/Xray-core/main/confloader"
)

func init() {
	common.Must(core.RegisterConfigLoader(&core.ConfigFormat{
		Name:      "YAML",
		Extension: []string{"yaml", "yml"},
		Loader: func(input interface{}) (*core.Config, error) {
			switch v := input.(type) {
			case cmdarg.Arg:
				cf := &conf.Config{}
				for i, arg := range v {
					newError("Reading config: ", arg).AtInfo().WriteToLog()
					r, err := confloader.LoadConfig(arg)
					if err != nil {
						return nil, newError("failed to read config: ", arg).Base(err)
					}
					c, err := serial.DecodeYAMLConfig(r)
					if err != nil {
						return nil, newError("failed to decode config: ", arg).Base(err)
					}
					if i == 0 {
						// This ensure even if the muti-json parser do not support a setting,
						// It is still respected automatically for the first configure file
						*cf = *c
						continue
					}
					cf.Override(c, arg)
				}
				return cf.Build()
			case io.Reader:
				return serial.LoadYAMLConfig(v)
			default:
				return nil, newError("unknow type")
			}
		},
	}))
}
