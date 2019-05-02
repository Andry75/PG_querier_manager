package queues

import "net"

type QueueDB interface {
	GetIP() net.IPAddr
	Insert(ip net.IPAddr) error
	SelectCountOfActive() (int, error)
	SelectNotActive() ([]net.IPAddr, error)
	SelectTheMostNotLoaded() error
	UpdateActiveJobsCount(ip net.IPAddr) error
}
