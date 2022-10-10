// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "hello/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	EchoEndpoint   endpoint.Endpoint
	HealthEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.HelloService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		EchoEndpoint:   MakeEchoEndpoint(s),
		HealthEndpoint: MakeHealthEndpoint(s),
	}
	for _, m := range mdw["Echo"] {
		eps.EchoEndpoint = m(eps.EchoEndpoint)
	}
	for _, m := range mdw["Health"] {
		eps.HealthEndpoint = m(eps.HealthEndpoint)
	}
	return eps
}
