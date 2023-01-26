package hello

import (
	"context"
	"github.com/sirupsen/logrus"
	"rest-template/pkg/service"
	"rest-template/pkg/service/endpoint"
)

// Request é a estrutura responsável pela interação da comunicação externa com o serviço
type Request struct {
	Request string
}

// Method método responsável por retornar o nome do endpoint; ação do endpoint
func (r Request) Method() string {
	return "hello"
}

// Response é a estrutura responsável pela resposta do serviço, composta pelo business.Preference
type Response struct {
	Data string `json:",inline"`
	Err  error  `json:"-"`
}

func (r Response) Failed() error {
	return r.Err
}

// New é uma função responsável pela criação do Endpoint
func New(svc service.Service) endpoint.Endpoint {
	return endpoint.EndpointFunc(func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*Request)

		ser := svc.Hello()
		logrus.Info("REQUEST: ", req.Request)

		return &Response{
			Data: "Hello",
			Err:  ser,
		}, nil
	})
}
