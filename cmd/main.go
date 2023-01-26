package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"os/signal"
	"rest-template/pkg/service/endpoint/createAccount"
	"rest-template/pkg/service/endpoint/createTransaction"
	"rest-template/pkg/service/endpoint/getAccount"
	"rest-template/pkg/service/endpoint/hello"
	"rest-template/pkg/transport"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
	"rest-template/config"
	"rest-template/config/connect"
	_ "rest-template/config/setup"
	"rest-template/pkg/service"
	"rest-template/pkg/service/endpoint"
	"rest-template/pkg/service/endpoint/middleware"
	"rest-template/pkg/transport/handler"
	httptransport "rest-template/pkg/transport/http"
)

const EXIT = syscall.Signal(177)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	env, err := config.LoadConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	mongoClient, err := connect.MongoDB(ctx, options.Client().ApplyURI(env.Mongo.ConnectionString))
	if err != nil {
		logrus.Fatal(err)
	}

	repo := service.NewRepository(mongoClient, env)
	svc := service.NewService(repo)

	var helloEndpoint endpoint.Endpoint
	{
		helloEndpoint = hello.New(svc)
		helloEndpoint = middleware.Limiter(rate.NewLimiter(rate.Every(time.Second), 100))(helloEndpoint)
		helloEndpoint = middleware.Logging()(helloEndpoint)
	}

	var createAccountEndpoint endpoint.Endpoint
	{
		createAccountEndpoint = createAccount.New(svc)
		createAccountEndpoint = middleware.Limiter(rate.NewLimiter(rate.Every(time.Second), 100))(createAccountEndpoint)
		createAccountEndpoint = middleware.Logging()(createAccountEndpoint)
	}

	var createTransactionEndpoint endpoint.Endpoint
	{
		createTransactionEndpoint = createTransaction.New(svc)
		createTransactionEndpoint = middleware.Limiter(rate.NewLimiter(rate.Every(time.Second), 100))(createTransactionEndpoint)
		createTransactionEndpoint = middleware.Logging()(createTransactionEndpoint)
	}

	var getAccountEndpoint endpoint.Endpoint
	{
		getAccountEndpoint = getAccount.New(svc)
		getAccountEndpoint = middleware.Limiter(rate.NewLimiter(rate.Every(time.Second), 100))(getAccountEndpoint)
		getAccountEndpoint = middleware.Logging()(getAccountEndpoint)
	}

	endpoints := &endpoint.Endpoints{
		Hello:             helloEndpoint,
		CreateAccount:     createAccountEndpoint,
		CreateTransaction: createTransactionEndpoint,
		GetAccount:        getAccountEndpoint,
	}

	echo := handler.NewHTTP(endpoints)

	transports := []transport.Transport{
		httptransport.New(echo, env.HTTP.Port),
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	for _, t := range transports {
		go func(t transport.Transport) {
			if err := t.Start(); err != nil {
				logrus.Error(err)
				sig <- EXIT
			}
		}(t)
	}

	<-sig

	for _, t := range transports {
		t.Stop()
	}
}
