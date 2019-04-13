package request_sender

import "github.com/Andry75/PG_querier_manager/http_interfaces"

func SendRequest(request http_interfaces.Request) http_interfaces.Response {
	return Response{
		httpStatus: 200,
		message:    "",
		payload:    "{\"username\":\"xyz\",\"password\":\"xyz\"}",
	}
}
