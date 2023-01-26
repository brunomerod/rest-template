package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"

	"rest-template/pkg/service"
	"rest-template/pkg/service/endpoint/middleware"
)

func httpErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code

		c.JSON(code, map[string]interface{}{
			"err": err.Error(),
		})
		return
	}

	switch err {
	case middleware.ErrLimited:
		c.JSON(http.StatusTooManyRequests, map[string]interface{}{
			"err": err.Error(),
		})
		return

	case service.ErrNotFound:
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"err": err.Error(),
		})
		return

	case service.ErrAlreadyExists:
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"err": err.Error(),
		})
		return

	case service.ErrAccountDoesntExist:
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"err": err.Error(),
		})
		return
	}

	c.JSON(code, map[string]interface{}{
		"err": http.StatusText(code),
	})
}
