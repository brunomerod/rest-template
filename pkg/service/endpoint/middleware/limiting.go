package middleware

import (
	"context"
	"errors"

	"golang.org/x/time/rate"
	"rest-template/pkg/service/endpoint"
)

var ErrLimited = errors.New("rate limit exceeded")

func Limiter(limit *rate.Limiter) endpoint.Middleware {
	return func(end endpoint.Endpoint) endpoint.Endpoint {
		return endpoint.EndpointFunc(func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !limit.Allow() {
				return nil, ErrLimited
			}

			return end.Execute(ctx, request)
		})
	}
}
