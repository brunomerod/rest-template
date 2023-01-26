package endpoint

//go:generate mockgen -destination endpoint_mock.go -package endpoint rest-template/pkg/service/endpoint Endpoint,Request,Response

import (
	"context"
)

type (
	Endpoint interface {
		Execute(ctx context.Context, request interface{}) (response interface{}, err error)
	}

	EndpointFunc func(ctx context.Context, request interface{}) (response interface{}, err error)

	Middleware func(Endpoint) Endpoint

	Failer interface {
		Failed() error
	}

	Meta interface {
		Method() string
	}

	// Request mock for testing
	Request interface {
		Meta
	}

	// Response mock for testing
	Response interface {
		Failer
	}
)

func (endpoint EndpointFunc) Execute(ctx context.Context, request interface{}) (response interface{}, err error) {
	return endpoint(ctx, request)
}

type Method struct {
	Method string `json:"method"`
}

type Endpoints struct {
	Hello             Endpoint
	CreateAccount     Endpoint
	CreateTransaction Endpoint
	GetAccount        Endpoint
}
