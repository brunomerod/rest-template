package service

//go:generate mockgen -destination repository_mock.go -package service rest-template/pkg/service Repository

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"rest-template/config"
)

type repository struct {
	client      *mongo.Client
	environment *config.Env
}

type Repository interface {
	InsertAccount(ctx context.Context, account Account) error
	InsertTransaction(ctx context.Context, transaction Transaction) error
	GetAccounts(ctx context.Context) ([]Account, error)
	GetAccountById(ctx context.Context, id primitive.ObjectID) ([]Account, error)
}

const (
	accountsCollection     string = "accounts"
	transactionsCollection string = "transactions"
)

func (r *repository) accountsRepository() *mongo.Collection {
	opts := &options.CollectionOptions{
		ReadPreference: readpref.SecondaryPreferred(),
	}
	return r.client.Database(r.environment.Mongo.Database).Collection(accountsCollection, opts)
}

func (r *repository) transactionsRepository() *mongo.Collection {
	opts := &options.CollectionOptions{
		ReadPreference: readpref.SecondaryPreferred(),
	}
	return r.client.Database(r.environment.Mongo.Database).Collection(transactionsCollection, opts)
}

func NewRepository(client *mongo.Client, env *config.Env) Repository {
	return &repository{client, env}
}

func (r *repository) InsertAccount(ctx context.Context, account Account) error {
	if r.checkAccountByDocumentNumber(ctx, account.DocumentNumber) {
		return ErrAlreadyExists
	}

	_, err := r.accountsRepository().InsertOne(ctx, account)

	return err
}

func (r *repository) GetAccounts(ctx context.Context) ([]Account, error) {
	var accounts []Account

	cursor, err := r.accountsRepository().Find(ctx, bson.D{{}})
	if err != nil {
		return accounts, err
	}

	for cursor.Next(ctx) {
		var i Account
		err := cursor.Decode(&i)
		if err != nil {
			return nil, errors.New("error on decoding database response")
		}

		accounts = append(accounts, i)
	}

	return accounts, nil
}

func (r *repository) GetAccountById(ctx context.Context, id primitive.ObjectID) ([]Account, error) {
	var account Account
	var accounts []Account
	err := r.accountsRepository().FindOne(ctx, bson.M{"_id": id}).Decode(&account)
	if err != nil {
		return nil, err
	}

	accounts = append(accounts, account)

	return accounts, nil
}

func (r *repository) InsertTransaction(ctx context.Context, transaction Transaction) error {
	if !r.checkAccountById(ctx, transaction.AccountId) {
		return ErrAccountDoesntExist
	}

	_, err := r.transactionsRepository().InsertOne(ctx, transaction)

	return err
}

func (r *repository) checkAccountByDocumentNumber(ctx context.Context, documentNumber string) bool {
	var account Account
	err := r.accountsRepository().FindOne(ctx, bson.M{"document_number": documentNumber}).Decode(&account)
	if err != nil {
		logrus.Error(err)
	}

	if account.AccountId.IsZero() {
		return false
	}

	return true
}

func (r *repository) checkAccountById(ctx context.Context, id primitive.ObjectID) bool {
	var account Account
	err := r.accountsRepository().FindOne(ctx, bson.M{"_id": id}).Decode(&account)
	if err != nil {
		logrus.Error(err)
	}

	if account.AccountId.IsZero() {
		return false
	}

	return true
}
