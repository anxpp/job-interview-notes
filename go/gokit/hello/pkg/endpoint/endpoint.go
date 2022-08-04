package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "hello/pkg/service"
)

// EchoRequest collects the request parameters for the Echo method.
type EchoRequest struct {
	Src string `json:"src"`
}

// EchoResponse collects the response parameters for the Echo method.
type EchoResponse struct {
	S0 string `json:"s0"`
}

// MakeEchoEndpoint returns an endpoint that invokes Echo on the service.
func MakeEchoEndpoint(s service.HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(EchoRequest)
		s0 := s.Echo(ctx, req.Src)
		return EchoResponse{S0: s0}, nil
	}
}

// HealthRequest collects the request parameters for the Health method.
type HealthRequest struct{}

// HealthResponse collects the response parameters for the Health method.
type HealthResponse struct {
	S0 string `json:"s0"`
}

// MakeHealthEndpoint returns an endpoint that invokes Health on the service.
func MakeHealthEndpoint(s service.HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		s0 := s.Health(ctx)
		return HealthResponse{S0: s0}, nil
	}
}

// Echo implements Service. Primarily useful in a client.
func (e Endpoints) Echo(ctx context.Context, src string) (s0 string) {
	request := EchoRequest{Src: src}
	response, err := e.EchoEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(EchoResponse).S0
}

// Health implements Service. Primarily useful in a client.
func (e Endpoints) Health(ctx context.Context) (s0 string) {
	request := HealthRequest{}
	response, err := e.HealthEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(HealthResponse).S0
}
