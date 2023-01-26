package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockgen -destination service_mock.go -package service rest-template/pkg/service Service

type Service interface {
	Hello() error
	CreateAccount(ctx context.Context, account Account) error
	CreateTransaction(ctx context.Context, transaction Transaction) error
	GetAccount(ctx context.Context, accountId primitive.ObjectID) ([]Account, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Hello() error {
	return nil
}

func (s *service) CreateAccount(ctx context.Context, account Account) error {
	err := s.repo.InsertAccount(ctx, account)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) CreateTransaction(ctx context.Context, transaction Transaction) error {
	a, err := GetTransactionValue(transaction.OperationTypeId, transaction.Amount)
	if err != nil {
		return err
	}

	transaction.Amount = a

	err = s.repo.InsertTransaction(ctx, transaction)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetAccount(ctx context.Context, accountId primitive.ObjectID) ([]Account, error) {
	if accountId.IsZero() {
		return s.repo.GetAccounts(ctx)
	}

	return s.repo.GetAccountById(ctx, accountId)
}
