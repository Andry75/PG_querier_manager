package requests_router

import "net"

type Request struct {
	ipAddress    net.IPAddr
	payload      string
	endpointName string
	method       string
}

func (r Request) GetIpdAddress() net.IPAddr {
	return r.ipAddress
}

func (r Request) GetEndpointName() string {
	return r.endpointName
}

func (r Request) GetPayload() string {
	return r.payload
}

func (r Request) GetMethod() string {
	return r.method
}
