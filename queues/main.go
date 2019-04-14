package queues

import "net"

func FindTheMostNonOverloadedNode() net.IPAddr {
	return net.IPAddr{IP: net.ParseIP("10.10.1.1")}
}

func ReleaseNode(ip net.IPAddr) {
	println("Hi I'm queue")
}

func CountOfActiveNodes() int {
	return 0
}

func AddNewNodeToTheQueue(ip net.IPAddr) {

}

func FindNotActiveNodes() []net.IPAddr {
	return []net.IPAddr{}
}
