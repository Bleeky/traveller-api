package main

import (
	"traveller-api/travelservice"
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
)

func main() {
	restful.DefaultResponseContentType(restful.MIME_JSON)
	restful.DefaultRequestContentType(restful.MIME_JSON)
	restful.Add(travelservice.New())
	server := &http.Server{Addr: ":8080"}
	log.Fatal(server.ListenAndServe())
}
