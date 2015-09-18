package main

import (
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
	"traveller-api/placeservice"
	"traveller-api/travelservice"
)

func main() {
	restful.DefaultResponseContentType(restful.MIME_JSON)
	restful.DefaultRequestContentType(restful.MIME_JSON)
	restful.Add(travelservice.New())
	restful.Add(placeservice.New())
	server := &http.Server{Addr: ":8080"}
	log.Fatal(server.ListenAndServe())
}
