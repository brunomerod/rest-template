package http

import (
	"context"
	"github.com/labstack/echo/v4"

	"rest-template/pkg/service/endpoint"
)

func EncodeJSON(statusCode int) EncodeFunc {
	return func(_ context.Context, c echo.Context, response interface{}) error {
		return c.JSON(statusCode, response)
	}
}

func Serve(end endpoint.Endpoint, dec DecodeFunc, enc EncodeFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		request, err := dec(ctx, c)
		if err != nil {
			return err
		}

		response, err := end.Execute(ctx, request)
		if err != nil {
			return err
		}

		if fail, ok := response.(endpoint.Failer); ok {
			if err = fail.Failed(); err != nil {
				return err
			}
		}

		return enc(ctx, c, response)
	}
}
