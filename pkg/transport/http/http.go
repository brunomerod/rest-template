package http

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"time"

	"github.com/sirupsen/logrus"
	"rest-template/pkg/transport"
)

type httpTransport struct {
	router *echo.Echo
	port   int
}

func New(router *echo.Echo, port int) transport.Transport {
	return &httpTransport{
		router: router,
		port:   port,
	}
}

func (transport *httpTransport) Start() error {
	logrus.Infof("HTTP server started on :%d", transport.port)
	return transport.router.Start(fmt.Sprintf(":%d", transport.port))
}

func (transport *httpTransport) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	transport.router.Shutdown(ctx)
}
