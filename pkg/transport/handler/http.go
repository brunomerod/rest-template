package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"rest-template/pkg/service"
	"rest-template/pkg/service/endpoint"
	"rest-template/pkg/service/endpoint/createAccount"
	"rest-template/pkg/service/endpoint/createTransaction"
	"rest-template/pkg/service/endpoint/getAccount"
	"rest-template/pkg/service/endpoint/hello"
	httphandler "rest-template/pkg/transport/handler/http"
	"time"
)

func NewHTTP(endpoints *endpoint.Endpoints) *echo.Echo {
	e := echo.New()
	e.HTTPErrorHandler = httpErrorHandler
	e.HideBanner = false
	e.HidePort = false

	e.GET("/hello",
		httphandler.Serve(
			endpoints.Hello,
			decodeHello,
			httphandler.EncodeJSON(http.StatusOK)))

	e.POST("/accounts",
		httphandler.Serve(
			endpoints.CreateAccount,
			decodeCreateAccount,
			httphandler.EncodeJSON(http.StatusNoContent)))

	e.POST("/transactions",
		httphandler.Serve(
			endpoints.CreateTransaction,
			decodeCreateTransaction,
			httphandler.EncodeJSON(http.StatusNoContent)))

	e.GET("/accounts",
		httphandler.Serve(
			endpoints.GetAccount,
			decodeGetAccount,
			httphandler.EncodeJSON(http.StatusOK)))

	e.GET("/accounts/:accountId",
		httphandler.Serve(
			endpoints.GetAccount,
			decodeGetAccount,
			httphandler.EncodeJSON(http.StatusOK)))

	return e
}

func decodeHello(ctx context.Context, c echo.Context) (req interface{}, err error) {
	req = &hello.Request{
		Request: "hello",
	}

	return req, nil
}

func decodeCreateAccount(ctx context.Context, c echo.Context) (req interface{}, err error) {
	var body service.Account

	err = json.NewDecoder(c.Request().Body).Decode(&body)
	if err != nil {
		return nil, errors.New("unable to decode body")
	}

	req = &createAccount.Request{
		Account: body,
	}

	return
}

func decodeCreateTransaction(ctx context.Context, c echo.Context) (req interface{}, err error) {
	var body service.Transaction

	err = json.NewDecoder(c.Request().Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	body.EventDate = time.Now().String()

	req = &createTransaction.Request{
		Transaction: body,
	}

	return
}

func decodeGetAccount(ctx context.Context, c echo.Context) (req interface{}, err error) {
	if c.Param("accountId") == "" {
		var nullId primitive.ObjectID
		req = &getAccount.Request{
			AccountId: nullId,
		}

		return
	}

	accountId, err := primitive.ObjectIDFromHex(c.Param("accountId"))
	if err != nil {
		return nil, err
	}

	req = &getAccount.Request{
		AccountId: accountId,
	}

	return
}
