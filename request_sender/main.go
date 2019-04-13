package request_sender

import (
	"bytes"
	"github.com/Andry75/PG_querier_manager/http_interfaces"
	"io/ioutil"
	"net/http"
	"time"
)

func SendRequest(request http_interfaces.Request) http_interfaces.Response {
	req, err := createRequest(request)
	if err != nil {
		println(err.Error())
		return Response{
			httpStatus: http.StatusInternalServerError,
			message:    "500 Internal Server Error",
			payload:    "",
		}
	}

	resp, err := dispatchRequest(req)
	if err != nil {
		println(err.Error())
		return Response{
			httpStatus: http.StatusInternalServerError,
			message:    "500 Internal Server Error",
			payload:    "",
		}
	}

	return resp
}

func generateURL(request http_interfaces.Request) string {
	ip := request.GetIpdAddress()
	return "http://" + ip.String() + request.GetEndpointName()
}

func createRequest(request http_interfaces.Request) (*http.Request, error) {
	newRequest, err := http.NewRequest(request.GetMethod(),
		generateURL(request),
		bytes.NewBuffer([]byte(request.GetPayload())))
	if err != nil {
		return nil, err
	}
	setHeaders(newRequest)
	return newRequest, nil
}

func setHeaders(r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
}

func dispatchRequest(r *http.Request) (http_interfaces.Response, error) {
	client := &http.Client{Timeout: time.Second * 50}
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	payload, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return Response{
		httpStatus: resp.StatusCode,
		message:    resp.Status,
		payload:    string(payload),
	}, nil
}
