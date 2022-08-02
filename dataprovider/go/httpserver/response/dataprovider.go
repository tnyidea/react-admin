package response

import (
	httpresponse "github.com/tnyidea/go-httpserver/response"
	"github.com/tnyidea/typeutils"
	"log"
	"net/http"
	"runtime/debug"
	"strconv"
)

type DataProviderResponseData struct {
	Status int
	Items  any

	TotalCount int
}

func WriteDataProviderResponseData(w http.ResponseWriter, r *http.Request, response DataProviderResponseData) {
	w.Header().Set("X-Total-Count", strconv.FormatInt(int64(response.TotalCount), 10))

	statusCode := typeutils.IntDefault(response.Status, http.StatusOK)

	w = httpresponse.WithHeaderNoCache(w)
	w = httpresponse.WithHeaderContentType(w, httpresponse.HeaderContentTypeApplicationJson)
	w.Header().Set("access-control-expose-headers", "X-Total-Count")

	err := httpresponse.WriteJsonResponse(w, r, statusCode, response.Items)
	if err != nil {
		log.Println("error: Error calling response.WriteJsonResponse():", err)
		debug.PrintStack()
		log.Fatal("fatal: Exiting")
	}
}
