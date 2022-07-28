package httpserver

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (p *HttpServer) NewRouter() *mux.Router {
	// Configure Router
	router := mux.NewRouter()

	// Diagnostic Routes
	router.HandleFunc("/healthcheck", HealthCheck).Methods(http.MethodGet)
	router.HandleFunc("/helloworld", HelloWorld).Methods(http.MethodGet)

	// Additional Routers
	p.AddRouterUsersV1(router)

	return router
}

func (p *HttpServer) AddRouterUsersV1(router *mux.Router) {
	router.HandleFunc("/api/v1/users", p.CreateUserV1).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/users", p.FindAllUsersV1).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/users/{uuid}", p.FindUserByUUIDV1).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/users/{uuid}", p.UpdateUserV1).Methods(http.MethodPut)
	router.HandleFunc("/api/v1/users/{uuid}", p.DeleteUserByUUIDV1).Methods(http.MethodDelete)
}
