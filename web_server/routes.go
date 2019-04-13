package web_server

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter()
	router.PathPrefix("/").HandlerFunc(Any)

	return router
}
