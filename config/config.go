package config

import (
	"errors"
	"github.com/joho/godotenv"
	logger "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type Env struct {
	HTTP  HTTP
	Mongo Mongo
}

type HTTP struct {
	Port int
}

type Mongo struct {
	ConnectionString string
	Database         string
}

var (
	ErrEnvMongoURL       = errors.New("env var isn't set: MONGO_URL")
	ErrEnvDatabase       = errors.New("env var isn't set: DATABASE")
	ErrConvertEnvApiPort = errors.New("error converting PORT to int")
)

func LoadConfig() (*Env, error) {
	logger.Info("loading environment variables")

	gaeEnv := os.Getenv("GAE_ENV")
	if gaeEnv == "" {
		err := godotenv.Load("./config/.env")
		if err != nil {
			logger.WithError(err).Fatal("failed to load config: error loading .env file")
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		logger.Infof("defaulting to port %s", port)
	}

	apiPort, err := strconv.Atoi(port)
	if err != nil {
		logger.WithError(ErrConvertEnvApiPort)
		return nil, errors.New(ErrConvertEnvApiPort.Error())
	}

	connectionStringDB, ok := os.LookupEnv("MONGO_URL")
	if !ok {
		logger.WithError(ErrEnvMongoURL)
		return nil, errors.New(ErrEnvMongoURL.Error())
	}

	database, ok := os.LookupEnv("DATABASE")
	if !ok {
		logger.WithError(ErrEnvDatabase)
		return nil, errors.New(ErrEnvDatabase.Error())
	}

	return &Env{
		HTTP: HTTP{Port: apiPort},
		Mongo: Mongo{
			ConnectionString: connectionStringDB,
			Database:         database,
		},
	}, nil
}
