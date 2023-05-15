package router

import (
	"log"
	"net/http"
	"url-shorten-service/pkg/handler"
	"url-shorten-service/pkg/utils"
)

func HandlRoutes() {
	http.HandleFunc("/shortner/encode", handler.EnocdeURL)
	http.HandleFunc("/shortner/decode", handler.DecodeURL)
	log.Fatal(http.ListenAndServe(utils.PORT, nil))
}
