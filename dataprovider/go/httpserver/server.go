package httpserver

import (
	"log"
	"net/http"
	"runtime/debug"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/tnyidea/react-admin-dataprovider/go/data"
	"github.com/tnyidea/typeutils"
)

// TODO set this up in Vault?

type HttpServer struct {
	Port int
	DB   data.DB
}

func NewHttpServer(port int) (HttpServer, error) {
	db, err := data.NewUserDatabase("../data/us-500.json")
	if err != nil {
		return HttpServer{}, err
	}

	return HttpServer{
		Port: typeutils.IntDefault(port, 80),
		DB:   db,
	}, nil
}

func (p *HttpServer) ListenAndServe() error {
	// Start the Server
	log.Printf("HTTP Server Listening on port %d...", p.Port)

	return http.ListenAndServe(":"+strconv.FormatInt(int64(p.Port), 10), handlers.CORS()(p.NewRouter()))
}

func (p *HttpServer) Defer() {
	// Capture a slice of errors for multiple defer actions
	var errs []error

	// Close the database
	err := p.DB.Close()
	if err != nil {
		debug.PrintStack()
		errs = append(errs, err)
	}

	// Handle errors
	if errs != nil {
		log.Fatal("fatal: Error while exiting:", errs)
	}
}
