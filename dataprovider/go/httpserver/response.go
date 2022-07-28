package httpserver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/tnyidea/typeutils"
	"reflect"
	"runtime/debug"
	"time"
)

type Response struct {
	ApiVersion *string        `json:"version,omitempty"`
	Id         *string        `json:"id,omitempty"`
	Method     *string        `json:"method,omitempty"`
	Data       *DataResponse  `json:"data,omitempty"`
	Error      *ErrorResponse `json:"error,omitempty"`
}

type ErrorResponse struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

type DataResponse struct {
	Code    *int        `json:"code,omitempty"`
	Message *string     `json:"message,omitempty"`
	Items   interface{} `json:"items,omitempty"`
}

// WriteResponse
// TODO Separate request validation from here
// TODO Separate header configuration from here
func WriteResponse(w http.ResponseWriter, r *http.Request, response Response) {
	// Validate request
	requestId := r.Context().Value(RequestContextId).(uuid.UUID).String()
	if requestId == "" {
		r, requestId = NewRequestContextId(r)
		log.SetPrefix(requestId + ": ")
		log.Printf("error: Must initialize RequestContextId. Overriding value to %v", requestId)
		log.Println("calling: WriteErrorResponse()")
		debug.PrintStack()
		WriteErrorResponse(w, r)
		return
	}
	log.SetPrefix(requestId + ": ")

	apiVersion := r.Context().Value(RequestContextApiVersion).(string)
	if apiVersion == "" {
		log.Printf("error: Must initialize RequestContextApiVersion")
		log.Println("calling: WriteErrorResponse()")
		debug.PrintStack()
		WriteErrorResponse(w, r)
		return
	}

	apiMethod := r.Context().Value(RequestContextApiMethod).(string)
	if apiMethod == "" {
		log.Printf("error: Must initialize RequestContextApiMethod")
		log.Println("calling: WriteErrorResponse()")
		debug.PrintStack()
		WriteErrorResponse(w, r)
		return
	}

	// Validate response
	var statusCode int
	if response.Data != nil && response.Error != nil ||
		response.Data == nil && response.Error == nil {
		log.Println("error: Must specify one of Response.Data or Response.Error")
		log.Println("calling: WriteErrorResponse()")
		debug.PrintStack()
		WriteErrorResponse(w, r)
		return
	}
	if response.Data != nil {
		if response.Data.Items != nil {
			if reflect.TypeOf(response.Data.Items).Kind() != reflect.Slice {
				log.Println("error: Response.Data.Items must be of kind Slice")
				log.Println("calling: WriteErrorResponse()")
				debug.PrintStack()
				WriteErrorResponse(w, r)
				return
			}
		}
		if response.Data.Code == nil {
			response.Data.Code = typeutils.IntPtr(http.StatusOK)
		}
		statusCode = *response.Data.Code
	}
	if response.Error != nil {
		if response.Error.Code == nil {
			response.Error.Code = typeutils.IntPtr(http.StatusInternalServerError)
		}
		statusCode = *response.Error.Code
	}

	// Set header no-cache
	w.Header().Set("Cache-Control", "must-revalidate, no-cache, no-store, no-transform, max-age=0")
	w.Header().Set("Expires", time.Unix(0, 0).Format(http.TimeFormat))
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("X-Accel-Expires", "0")

	// Set header content-type
	w.Header().Set("Content-Type", "application/json")

	// Set response data
	response.Id = typeutils.StringPtr(requestId)
	response.ApiVersion = typeutils.StringPtr(apiVersion)
	response.Method = typeutils.StringPtr(apiMethod)

	b, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Printf("error: Error calling json.MarshalIndent(): %v", err)
		debug.PrintStack()
		log.Fatal("fatal: Exiting")
	}

	w.WriteHeader(statusCode)
	_, err = w.Write(b)
	if err != nil {
		log.Printf("error: Error calling w.Write(): %v", err)
		debug.PrintStack()
		log.Fatal("fatal: Exiting")
	}
}

func WriteEmptyResponse(w http.ResponseWriter, r *http.Request) {
	WriteResponse(w, r, Response{
		Data: &DataResponse{
			Code:    typeutils.IntPtr(http.StatusOK),
			Message: typeutils.StringPtr(http.StatusText(http.StatusOK)),
		},
	})
}

func WriteErrorResponse(w http.ResponseWriter, r *http.Request) {
	WriteResponse(w, r, Response{
		Error: &ErrorResponse{
			Code:    typeutils.IntPtr(http.StatusInternalServerError),
			Message: typeutils.StringPtr(http.StatusText(http.StatusInternalServerError)),
		},
	})
}
