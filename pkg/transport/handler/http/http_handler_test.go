package http

import (
	"bytes"
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"rest-template/pkg/service/endpoint"
)

func TestServe(t *testing.T) {
	require := require.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEndpoint := endpoint.NewMockEndpoint(ctrl)
	mockEndpoint.EXPECT().Execute(gomock.Any(), gomock.Any()).Return(nil, nil)

	decode := func(ctx context.Context, c echo.Context) (interface{}, error) {
		return nil, nil
	}

	encode := func(ctx context.Context, c echo.Context, response interface{}) error {
		return nil
	}

	handle := Serve(mockEndpoint, decode, encode)

	request := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()
	c := e.NewContext(request, nil)

	err := handle(c)
	require.NoError(err)
}

func TestServeDecodeError(t *testing.T) {
	require := require.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEndpoint := endpoint.NewMockEndpoint(ctrl)

	decode := func(ctx context.Context, c echo.Context) (interface{}, error) {
		return nil, errors.New("Erro no decode")
	}

	encode := func(ctx context.Context, c echo.Context, response interface{}) error {
		return nil
	}

	handle := Serve(mockEndpoint, decode, encode)

	request := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()
	c := e.NewContext(request, nil)

	err := handle(c)
	require.Error(err)
}

func TestServeEndpointError(t *testing.T) {
	require := require.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEndpoint := endpoint.NewMockEndpoint(ctrl)
	mockEndpoint.EXPECT().Execute(gomock.Any(), gomock.Any()).Return(nil, errors.New("Erro de ep"))

	decode := func(ctx context.Context, c echo.Context) (interface{}, error) {
		return nil, nil
	}

	encode := func(ctx context.Context, c echo.Context, response interface{}) error {
		return nil
	}

	handle := Serve(mockEndpoint, decode, encode)

	request := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()
	c := e.NewContext(request, nil)

	err := handle(c)
	require.Error(err)
}

func TestServeFailedError(t *testing.T) {
	require := require.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockResponse := endpoint.NewMockResponse(ctrl)
	mockResponse.EXPECT().Failed().Return(errors.New("Failed"))

	mockEndpoint := endpoint.NewMockEndpoint(ctrl)
	mockEndpoint.EXPECT().Execute(gomock.Any(), gomock.Any()).Return(mockResponse, nil)

	decode := func(ctx context.Context, c echo.Context) (interface{}, error) {
		return nil, nil
	}

	encode := func(ctx context.Context, c echo.Context, response interface{}) error {
		return nil
	}

	handle := Serve(mockEndpoint, decode, encode)

	request := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()
	c := e.NewContext(request, nil)

	err := handle(c)
	require.Error(err)
}

func TestEncodeJSON(t *testing.T) {
	enc := EncodeJSON(http.StatusOK)

	require := require.New(t)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(request, recorder)

	err := enc(context.TODO(), c, map[string]interface{}{
		"err": "err",
	})

	require.NoError(err)

	response := recorder.Result()
	var body bytes.Buffer
	_, err = body.ReadFrom(response.Body)

	require.NoError(err)
	require.JSONEq(`{"err":"err"}`, body.String())
}
