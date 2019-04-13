package web_server

import (
	"github.com/Andry75/PG_querier_manager/http_interfaces"
	"github.com/Andry75/PG_querier_manager/requests_router"
	"io/ioutil"
	"net/http"
)

func Any(w http.ResponseWriter, r *http.Request) {
	query := getQuery(w, r)
	response := requests_router.Route(query)
	sendResponse(response, w)
}

func sendResponse(response http_interfaces.Response, w http.ResponseWriter) {
	w.WriteHeader(response.GetHttpStatus())
	if response.MessageEmpty() {
		w.Write([]byte(response.GetPayload()))
	} else {
		w.Write([]byte(response.GetMassage()))
	}
}

func getQuery(w http.ResponseWriter, r *http.Request) Query {
	payload, err := getPayload(r)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("422 (Unprocessable Entity)"))
	}
	endpointName := getEndPointName(r)
	return Query{
		endpointName: endpointName,
		payload:      payload,
	}
}

func getPayload(r *http.Request) (string, error) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func getEndPointName(r *http.Request) string {
	return r.RequestURI
}
