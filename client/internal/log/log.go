package log

import (
	"client/internal/apperrors"
	"os"

	"github.com/sirupsen/logrus"
)

const logFile = "app.log"

func NewLogAndSetLevel(logLevel string) (*logrus.Logger, error) {
	log := logrus.New()
	loggerLevel, err := logrus.ParseLevel(logLevel)
	if err != nil {
		appErr := apperrors.NewLoggerErr.AppendMessage(err)
		return nil, appErr
	}

	log.SetLevel(loggerLevel)
	log.SetReportCaller(true)

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		appErr := apperrors.NewLoggerErr.AppendMessage(err)
		return nil, appErr
	}

	//log.SetOutput(os.Stdout)
	log.SetOutput(file)
	log.Info("Logger has been configurated")
	return log, nil
}

func SetLevel(log *logrus.Logger, logLevel string) error {
	loggerLevel, err := logrus.ParseLevel(logLevel)
	if err != nil {
		appErr := apperrors.NewLoggerErr.AppendMessage(err)
		return appErr
	}

	log.SetLevel(loggerLevel)
	log.Info("logger level has been configurated")
	return nil
}
