package http

//go:generate mockgen -destination encoder_decoder_mock.go -package http rest-template/pkg/transport/handler/http Decoder,Encoder

import (
	"context"
	"github.com/labstack/echo/v4"
)

type Decoder interface {
	Decode(context.Context, echo.Context) (interface{}, error)
}

type Encoder interface {
	Encode(context.Context, echo.Context, interface{}) error
}

type DecodeFunc func(context.Context, echo.Context) (interface{}, error)
type EncodeFunc func(context.Context, echo.Context, interface{}) error
