package connect

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDB(ctx context.Context, opts ...*options.ClientOptions) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, opts...)
	if err != nil {
		logrus.WithError(err).Info("error on connect mongo db")
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		logrus.WithError(err).Info("error on ping in mongo db")
		return nil, err
	}

	return client, nil
}
