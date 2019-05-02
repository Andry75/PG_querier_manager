package queues

import (
	"github.com/Andry75/PG_querier_manager/pg_adapter"
	"log"
	"net"
)

func FindTheMostNonOverloadedNode() net.IPAddr {
	q := getQueue()
	err := q.SelectTheMostNotLoaded()
	if err != nil {
		log.Println(err)
		return net.IPAddr{}
	}
	return q.GetIP()
}

func ReleaseNode(ip net.IPAddr) {
	q := getQueue()
	err := q.UpdateActiveJobsCount(ip)
	if err != nil {
		log.Println(err)
	}
}

func CountOfActiveNodes() int {
	q := getQueue()
	count, err := q.SelectCountOfActive()
	if err != nil {
		log.Println(err)
		return 0
	}
	return count
}

func AddNewNodeToTheQueue(ip net.IPAddr) {
	q := getQueue()
	err := q.Insert(ip)
	if err != nil {
		log.Println(err)
	}
}

func FindNotActiveNodes() []net.IPAddr {
	q := getQueue()
	ips, err := q.SelectNotActive()
	if err != nil {
		log.Println(err)
	}

	return ips
}

func getQueue() QueueDB {
	q := pg_adapter.Queues{}
	return &q
}
