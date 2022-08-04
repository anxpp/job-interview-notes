package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(HelloService) HelloService

type loggingMiddleware struct {
	logger log.Logger
	next   HelloService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a HelloService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next HelloService) HelloService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Echo(ctx context.Context, src string) (s0 string) {
	defer func() {
		l.logger.Log("method", "Echo", "src", src, "s0", s0)
	}()
	return l.next.Echo(ctx, src)
}
func (l loggingMiddleware) Health(ctx context.Context) (s0 string) {
	defer func() {
		l.logger.Log("method", "Health", "s0", s0)
	}()
	return l.next.Health(ctx)
}
