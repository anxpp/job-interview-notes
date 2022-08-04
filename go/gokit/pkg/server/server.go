package server

import (
	"GoNotes/go/gokit/pkg/endpoint"
	"GoNotes/go/gokit/pkg/transport"
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer 是一个很好的服务器
func NewHTTPServer(ctx context.Context, endpoints endpoint.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware) // @请参阅 https://stackoverflow.com/a/51456342

	r.Methods("GET").Path("/status").Handler(httptransport.NewServer(
		endpoints.StatusEndpoint,
		transport.DecodeStatusRequest,
		transport.EncodeResponse,
	))

	r.Methods("GET").Path("/get").Handler(httptransport.NewServer(
		endpoints.GetEndpoint,
		transport.DecodeGetRequest,
		transport.EncodeResponse,
	))

	r.Methods("POST").Path("/validate").Handler(httptransport.NewServer(
		endpoints.ValidateEndpoint,
		transport.DecodeValidateRequest,
		transport.EncodeResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
