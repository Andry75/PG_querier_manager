package web_server

import (
	"github.com/Andry75/PG_querier_manager/config_loader"
	"log"
	"net/http"
)

func Start() {
	config := getConfigs()
	log.Fatal(http.ListenAndServe(config.GetWebServerPort(), NewRouter()))
}

func getConfigs() WebServerConfig {
	return config_loader.Load()
}
