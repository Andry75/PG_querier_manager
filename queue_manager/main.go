package queue_manager

import "net"

func GetAvailableNode() net.IPAddr {
	return net.IPAddr{IP: net.ParseIP("10.10.1.1")}
}

func ReleaseNode(addr net.IPAddr) {
	println("Hi I'm queue")
}
