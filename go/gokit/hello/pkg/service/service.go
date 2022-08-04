package service

import "context"

// HelloService describes the service.
type HelloService interface {
	Echo(ctx context.Context, src string) string
	Health(ctx context.Context) string
}

type basicHelloService struct{}

func (b *basicHelloService) Echo(ctx context.Context, src string) (s0 string) {
	// TODO implement the business logic of Echo
	return s0
}
func (b *basicHelloService) Health(ctx context.Context) (s0 string) {
	// TODO implement the business logic of Health
	return s0
}

// NewBasicHelloService returns a naive, stateless implementation of HelloService.
func NewBasicHelloService() HelloService {
	return &basicHelloService{}
}

// New returns a HelloService with all of the expected middleware wired in.
func New(middleware []Middleware) HelloService {
	var svc HelloService = NewBasicHelloService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
