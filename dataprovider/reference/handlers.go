package reference

import (
	"encoding/json"
	"log"
	"net/http"
)

//
// Router Handler Functions
//

//
// Diagnostic Handlers
//

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&map[string]string{
		"result": http.StatusText(http.StatusOK),
	})
}

func logHeader(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, _ := json.MarshalIndent(r.Header, "", "    ")
		log.Println(string(b))
		handler(w, r)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&map[string]string{
		"message": "Hello World",
	})
}
