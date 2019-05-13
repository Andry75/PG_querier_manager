package instance_manager

import (
	"encoding/json"
	"fmt"
	"github.com/Andry75/PG_querier_manager/config_loader"
	"github.com/Andry75/PG_querier_manager/http_interfaces"
	"github.com/Andry75/PG_querier_manager/request_sender"
	"log"
	"net"
	"strconv"
)

func DeployNewNode() net.IPAddr {
	return sendRequestToDeployNewNode()
}

func WithdrawNode(ip net.IPAddr) {
	sendRequestToWithdrawNode(ip)
}

func sendRequestToDeployNewNode() net.IPAddr {
	configs := getConfigs()
	request := Request{
		ipAddress:    configs.GetInstancesMasterIpAddress(),
		endpointName: configs.GetInstancesMasterPort() + "/node",
		method:       "POST",
		payload:      "",
	}
	response := sendRequest(request)
	if response.GetHttpStatus() > 299 {
		log.Fatal(generateErrorMessage(net.IPAddr{}, response, "Deploy"))
		return net.IPAddr{}
	}

	responseData := Response{}
	err := json.Unmarshal([]byte(response.GetPayload()), &responseData)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Error while parsing response from Instances master:\n" +
			err.Error())
		return net.IPAddr{}
	}
	return responseData.GetIpAddress()
}

func sendRequestToWithdrawNode(ip net.IPAddr) {
	configs := getConfigs()
	request := Request{
		ipAddress:    configs.GetInstancesMasterIpAddress(),
		endpointName: configs.GetInstancesMasterPort() + "/node",
		method:       "DELETE",
		payload:      "{\"ip_address\" : \"" + ip.String() + "\"}",
	}
	response := sendRequest(request)
	if response.GetHttpStatus() > 299 {
		log.Println(generateErrorMessage(ip, response, "Withdraw"))
	}
}

func sendRequest(r Request) http_interfaces.Response {
	return request_sender.SendRequest(r)
}

func getConfigs() InstancesMasterConfig {
	return config_loader.Load()
}

func generateErrorMessage(nodeIp net.IPAddr, response http_interfaces.Response, operationName string) string {
	var nodeIpText = ""
	if len(nodeIp.String()) > 0 {
		nodeIpText = " node with IP: " + nodeIp.String()
	}
	return "An issue occurred while " + operationName + nodeIpText + ". Got the following:\n" +
		"HTTP Code: " + strconv.Itoa(response.GetHttpStatus()) + "\n" +
		"HTTP Message: " + response.GetMassage() + "\n"
}
