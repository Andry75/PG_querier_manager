package requests_router

import "github.com/Andry75/PG_querier_manager/http_interfaces"

func Route(q http_interfaces.Query) http_interfaces.Response {
	return Response{
		httpStatus: 200,
		message:    "",
		payload:    "{\"username\":\"xyz\",\"password\":\"xyz\"}",
	}
}
