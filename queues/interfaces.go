package queues

import "net"

type QueueReader interface {
	FindFreeNode() net.IPAddr
	CountOfActiveNodes() int
	FindRowWithoutJobs() net.IPAddr
}

type QueueWriter interface {
	ReduceJobsCount(ip net.IPAddr) error
	AddNewNode(ip net.IPAddr) error
	RemoveNode(ip net.IPAddr) error
}
