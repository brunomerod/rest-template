package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"rest-template/pkg/service/endpoint"
)

func Logging() endpoint.Middleware {
	return func(end endpoint.Endpoint) endpoint.Endpoint {
		return endpoint.EndpointFunc(func(ctx context.Context, request interface{}) (response interface{}, err error) {
			var requestMethod string

			if metadata, ok := request.(endpoint.Meta); ok {
				requestMethod = metadata.Method()
			}

			logrus.Infof("**starting new %v execution**", requestMethod)
			defer func(begin time.Time) {
				var erro error

				if failer, ok := response.(endpoint.Failer); ok {
					erro = failer.Failed()
				}

				if err != nil {
					erro = err
				}

				log := logrus.WithFields(logrus.Fields{
					"method": requestMethod,
					"error":  erro,
					"took":   fmt.Sprint(time.Since(begin)),
				})

				if erro != nil {
					log.Error()
				} else {
					log.Info()
				}
			}(time.Now())

			return end.Execute(ctx, request)
		})
	}
}
