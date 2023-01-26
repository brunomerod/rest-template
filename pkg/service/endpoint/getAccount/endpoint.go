package getAccount

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rest-template/pkg/service"
	"rest-template/pkg/service/endpoint"
)

// Request é a estrutura responsável pela interação da comunicação externa com o serviço
type Request struct {
	AccountId primitive.ObjectID
}

// Method método responsável por retornar o nome do endpoint; ação do endpoint
func (r Request) Method() string {
	return "get_account"
}

// Response é a estrutura responsável pela resposta do serviço, composta pelo business.Preference
type Response struct {
	Data []service.Account `json:",inline"`
	Err  error             `json:"-"`
}

func (r Response) Failed() error {
	return r.Err
}

// New é uma função responsável pela criação do Endpoint
func New(svc service.Service) endpoint.Endpoint {
	return endpoint.EndpointFunc(func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*Request)

		accounts, err := svc.GetAccount(ctx, req.AccountId)

		return &Response{
			Data: accounts,
			Err:  err,
		}, nil
	})
}
