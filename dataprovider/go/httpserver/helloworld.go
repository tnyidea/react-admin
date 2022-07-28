package httpserver

import (
	"log"
	"net/http"

	"github.com/tnyidea/typeutils"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	r, requestId := NewRequestContextId(r)
	log.SetPrefix(requestId + ": ")
	log.Println("=== Executing HelloWorld ===")

	r = SetApiRequestContext(r, "0", "get.helloWorld")

	WriteResponse(w, r, Response{
		Data: &DataResponse{
			Message: typeutils.StringPtr("Hello World"),
		},
	})

	log.Println("=== HelloWorld Execution Complete ===")
}
