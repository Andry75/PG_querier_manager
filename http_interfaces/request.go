package http_interfaces

import "net"

type Request interface {
	getIpdAddress() net.IPAddr
	getEndPointName() string
	getParams() string
}
