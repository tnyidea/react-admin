package router

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/tnyidea/go-httpserver/sample/userservice/httpserver/config"
	"github.com/tnyidea/go-sample-userdata/models"
	"github.com/tnyidea/react-admin/dataprovider/go/httpserver/response"
	"net/http"
)

func AddApiV1UsersRouter(router *mux.Router, ctx context.Context) *mux.Router {
	db := ctx.Value(config.UserServiceContextDatabase).(models.DB)

	router.HandleFunc("/api/v1/users", func(w http.ResponseWriter, r *http.Request) {

		// https://jsonplaceholder.typicode.com/users?_end=10&_order=ASC&_sort=id&_start=0

		users, err := db.FindAllUsers()
		if err != nil {

		}
		users = users[:10]

		response.WriteDataProviderResponseData(w, r, response.DataProviderResponseData{
			Status:     http.StatusOK,
			Items:      users,
			TotalCount: len(users),
		})
	}).Methods(http.MethodGet)

	return router
}
