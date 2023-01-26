package setup

import (
	"os"

	"github.com/sirupsen/logrus"
)

func setupLogger() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
}
