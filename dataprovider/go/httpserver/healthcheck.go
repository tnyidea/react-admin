package httpserver

import (
	"log"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	r, requestId := NewRequestContextId(r)
	log.SetPrefix(requestId + ": ")
	log.Println("=== Executing HealthCheck ===")

	r = SetApiRequestContext(r, "0", "get.healthCheck")

	WriteEmptyResponse(w, r)

	log.Println("=== HealthCheck Execution Complete ===")
}
