package instance_manager

import "net"

type InstancesMasterConfig interface {
	GetInstancesMasterIpAddress() net.IPAddr
	GetInstancesMasterPort() string
}
