package httpserver

import (
	"github.com/gorilla/mux"
	"github.com/tnyidea/react-admin-dataprovider/go/types"
	"log"
	"net/http"
	"runtime/debug"
)

func (p *HttpServer) CreateUserV1(w http.ResponseWriter, r *http.Request) {
	r, requestId := NewRequestContextId(r)
	log.SetPrefix(requestId + ": ")
	log.Println("=== Executing CreateUserV1 ===")

	r = SetApiRequestContext(r, "1", "post.createUser")

	err := p.DB.CreateUser(types.User{})
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		WriteErrorResponse(w, r)
		return
	}

	log.Println("=== CreateUserV1 Execution Complete ===")
}

func (p *HttpServer) FindAllUsersV1(w http.ResponseWriter, r *http.Request) {
	r, requestId := NewRequestContextId(r)
	log.SetPrefix(requestId + ": ")
	log.Println("=== Executing FindAllUsersV1 ===")

	r = SetApiRequestContext(r, "1", "get.findAllUsers")

	users, err := p.DB.FindAllUsers()
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		WriteErrorResponse(w, r)
		return
	}

	WriteResponse(w, r, Response{
		Data: &DataResponse{
			Items: users,
		},
	})

	log.Println("=== FindAllUsersV1 Execution Complete ===")
}

func (p *HttpServer) FindUserByUUIDV1(w http.ResponseWriter, r *http.Request) {
	r, requestId := NewRequestContextId(r)
	log.SetPrefix(requestId + ": ")
	log.Println("=== Executing FindUserByUUIDV1 ===")

	r = SetApiRequestContext(r, "1", "get.findUserByUUID")

	muxVars := mux.Vars(r)
	uuidString := muxVars["uuid"]

	var users []types.User
	user, err := p.DB.FindUserByUUID(uuidString)
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		WriteErrorResponse(w, r)
		return
	}
	users = append(users, user)

	WriteResponse(w, r, Response{
		Data: &DataResponse{
			Items: users,
		},
	})

	log.Println("=== FindUserByUUIDV1 Execution Complete ===")
}

func (p *HttpServer) UpdateUserV1(w http.ResponseWriter, r *http.Request) {
	r, requestId := NewRequestContextId(r)
	log.SetPrefix(requestId + ": ")
	log.Println("=== Executing CreateUserV1 ===")

	r = SetApiRequestContext(r, "1", "put.updateUser")

	err := p.DB.UpdateUser(types.User{})
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		WriteErrorResponse(w, r)
		return
	}

	log.Println("=== CreateUserV1 Execution Complete ===")
}

func (p *HttpServer) DeleteUserByUUIDV1(w http.ResponseWriter, r *http.Request) {
	r, requestId := NewRequestContextId(r)
	log.SetPrefix(requestId + ": ")
	log.Println("=== Executing DeleteUserV1 ===")

	r = SetApiRequestContext(r, "1", "delete.deleteUserByUUID")

	var uuidString string
	err := p.DB.DeleteUserByUUID(uuidString)
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		WriteErrorResponse(w, r)
		return
	}

	log.Println("=== DeleteUserV1 Execution Complete ===")
}
