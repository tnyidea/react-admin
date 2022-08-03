package router

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/tnyidea/go-httpserver/sample/userservice/httpserver/config"
	"github.com/tnyidea/go-sample-userdata/models"
	"github.com/tnyidea/react-admin/dataprovider/go/httpserver/endpoints"
	"github.com/tnyidea/react-admin/dataprovider/go/httpserver/response"
	"net/http"
)

func AddApiV1UsersRouter(router *mux.Router, ctx context.Context) *mux.Router {
	db := ctx.Value(config.UserServiceContextDatabase).(models.DB)

	router.HandleFunc("/api/v1/users", func(w http.ResponseWriter, r *http.Request) {
		response.WriteDataProviderResponseData(w, r, endpoints.FindAllUsersV1(r, db))
	}).Methods(http.MethodGet)

	return router
}
