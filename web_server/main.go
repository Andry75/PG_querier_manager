package web_server

import (
	"log"
	"net/http"
)

func Start() {
	log.Fatal(http.ListenAndServe(":8080", NewRouter()))
}
