package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "github.com/xxhanxx/Xray-core/app/dispatcher"
	_ "github.com/xxhanxx/Xray-core/app/proxyman/inbound"
	_ "github.com/xxhanxx/Xray-core/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "github.com/xxhanxx/Xray-core/app/commander"
	_ "github.com/xxhanxx/Xray-core/app/log/command"
	_ "github.com/xxhanxx/Xray-core/app/proxyman/command"
	_ "github.com/xxhanxx/Xray-core/app/stats/command"

	// Developer preview services
	_ "github.com/xxhanxx/Xray-core/app/observatory/command"

	// Other optional features.
	_ "github.com/xxhanxx/Xray-core/app/dns"
	_ "github.com/xxhanxx/Xray-core/app/dns/fakedns"
	_ "github.com/xxhanxx/Xray-core/app/log"
	_ "github.com/xxhanxx/Xray-core/app/metrics"
	_ "github.com/xxhanxx/Xray-core/app/policy"
	_ "github.com/xxhanxx/Xray-core/app/reverse"
	_ "github.com/xxhanxx/Xray-core/app/router"
	_ "github.com/xxhanxx/Xray-core/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "github.com/xxhanxx/Xray-core/transport/internet/tagged/taggedimpl"

	// Developer preview features
	_ "github.com/xxhanxx/Xray-core/app/observatory"

	// Inbound and outbound proxies.
	_ "github.com/xxhanxx/Xray-core/proxy/blackhole"
	_ "github.com/xxhanxx/Xray-core/proxy/dns"
	_ "github.com/xxhanxx/Xray-core/proxy/dokodemo"
	_ "github.com/xxhanxx/Xray-core/proxy/freedom"
	_ "github.com/xxhanxx/Xray-core/proxy/http"
	_ "github.com/xxhanxx/Xray-core/proxy/loopback"
	_ "github.com/xxhanxx/Xray-core/proxy/shadowsocks"
	_ "github.com/xxhanxx/Xray-core/proxy/socks"
	_ "github.com/xxhanxx/Xray-core/proxy/trojan"
	_ "github.com/xxhanxx/Xray-core/proxy/vless/inbound"
	_ "github.com/xxhanxx/Xray-core/proxy/vless/outbound"
	_ "github.com/xxhanxx/Xray-core/proxy/vmess/inbound"
	_ "github.com/xxhanxx/Xray-core/proxy/vmess/outbound"
	_ "github.com/xxhanxx/Xray-core/proxy/wireguard"

	// Transports
	_ "github.com/xxhanxx/Xray-core/transport/internet/domainsocket"
	_ "github.com/xxhanxx/Xray-core/transport/internet/grpc"
	_ "github.com/xxhanxx/Xray-core/transport/internet/http"
	_ "github.com/xxhanxx/Xray-core/transport/internet/kcp"
	_ "github.com/xxhanxx/Xray-core/transport/internet/quic"
	_ "github.com/xxhanxx/Xray-core/transport/internet/reality"
	_ "github.com/xxhanxx/Xray-core/transport/internet/tcp"
	_ "github.com/xxhanxx/Xray-core/transport/internet/tls"
	_ "github.com/xxhanxx/Xray-core/transport/internet/udp"
	_ "github.com/xxhanxx/Xray-core/transport/internet/websocket"

	// Transport headers
	_ "github.com/xxhanxx/Xray-core/transport/internet/headers/http"
	_ "github.com/xxhanxx/Xray-core/transport/internet/headers/noop"
	_ "github.com/xxhanxx/Xray-core/transport/internet/headers/srtp"
	_ "github.com/xxhanxx/Xray-core/transport/internet/headers/tls"
	_ "github.com/xxhanxx/Xray-core/transport/internet/headers/utp"
	_ "github.com/xxhanxx/Xray-core/transport/internet/headers/wechat"
	_ "github.com/xxhanxx/Xray-core/transport/internet/headers/wireguard"

	// JSON & TOML & YAML
	_ "github.com/xxhanxx/Xray-core/main/json"
	_ "github.com/xxhanxx/Xray-core/main/toml"
	_ "github.com/xxhanxx/Xray-core/main/yaml"

	// Load config from file or http(s)
	_ "github.com/xxhanxx/Xray-core/main/confloader/external"

	// Commands
	_ "github.com/xxhanxx/Xray-core/main/commands/all"
)
