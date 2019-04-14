package queue_manager

import (
	"github.com/Andry75/PG_querier_manager/instance_manager"
	"github.com/Andry75/PG_querier_manager/queues"
	"net"
	"time"
)

func GetAvailableNode() net.IPAddr {
	nodeIp := queues.FindTheMostNonOverloadedNode()
	if len(nodeIp.String()) == 0 && queues.CountOfActiveNodes() >= 49 {
		time.Sleep(10)
		return GetAvailableNode()
	}
	if len(nodeIp.String()) == 0 {
		newNodeIp := instance_manager.DeployNewNode()
		queues.AddNewNodeToTheQueue(newNodeIp)
		return newNodeIp
	} else {
		go withdrawNotActiveNodes()
		return nodeIp
	}
}

func ReleaseNode(addr net.IPAddr) {
	queues.ReleaseNode(addr)
}

func withdrawNotActiveNodes() {
	notActiveNodesIp := queues.FindNotActiveNodes()
	if len(notActiveNodesIp) > 0 {
		for _, nodeIp := range notActiveNodesIp {
			instance_manager.WithdrawNode(nodeIp)
		}
	}
}
