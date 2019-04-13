package http_interfaces

import "net"

type Request interface {
	GetIpdAddress() net.IPAddr
	GetEndpointName() string
	GetPayload() string
}
