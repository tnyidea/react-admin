package httpserver

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"time"
)

const RequestContextApiVersion = "requestContextApiVersion"
const RequestContextApiMethod = "requestContextApiMethod"
const RequestContextId = "requestContextId"
const RequestContextTimestamp = "requestContextTimestamp"

func NewRequestContextId(r *http.Request) (*http.Request, string) {
	uuidValue := uuid.New()
	ctx := context.WithValue(r.Context(), RequestContextId, uuidValue)
	ctx = context.WithValue(ctx, RequestContextTimestamp, time.Now())

	return r.Clone(ctx), uuidValue.String()
}

func SetApiRequestContext(r *http.Request, version string, method string) *http.Request {
	ctx := context.WithValue(r.Context(), RequestContextApiVersion, version)
	ctx = context.WithValue(ctx, RequestContextApiMethod, method)
	return r.Clone(ctx)
}
