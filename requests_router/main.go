package requests_router

import (
	"github.com/Andry75/PG_querier_manager/http_interfaces"
	"github.com/Andry75/PG_querier_manager/queue_manager"
	"github.com/Andry75/PG_querier_manager/request_sender"
	"net"
)

func Route(q http_interfaces.Query) http_interfaces.Response {
	request := generateRequest(q)
	defer queue_manager.ReleaseNode(request.GetIpdAddress())
	response := request_sender.SendRequest(request)
	return response
}

func generateRequest(query http_interfaces.Query) Request {
	return Request{
		endpointName: query.GetEndpointName(),
		payload:      query.GetPayload(),
		ipAddress:    getNodeIpAddress(),
		method:       query.GetMethod(),
	}
}

func getNodeIpAddress() net.IPAddr {
	return queue_manager.GetAvailableNode()
}
