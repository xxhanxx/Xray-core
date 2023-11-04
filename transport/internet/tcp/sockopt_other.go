//go:build !linux && !freebsd && !darwin
// +build !linux,!freebsd,!darwin

package tcp

import (
	"github.com/xxhanxx/Xray-core/common/net"
	"github.com/xxhanxx/Xray-core/transport/internet/stat"
)

func GetOriginalDestination(conn stat.Connection) (net.Destination, error) {
	return net.Destination{}, nil
}
