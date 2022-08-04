package http

import (
	"context"
	"encoding/json"
	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	endpoint "hello/pkg/endpoint"
	http1 "net/http"
)

// makeEchoHandler creates the handler logic
func makeEchoHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/echo").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.EchoEndpoint, decodeEchoRequest, encodeEchoResponse, options...)))
}

// decodeEchoRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeEchoRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.EchoRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeEchoResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeEchoResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeHealthHandler creates the handler logic
func makeHealthHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/health").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.HealthEndpoint, decodeHealthRequest, encodeHealthResponse, options...)))
}

// decodeHealthRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeHealthRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.HealthRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeHealthResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeHealthResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
