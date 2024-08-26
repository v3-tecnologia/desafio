package logger

import (
	"github.com/sirupsen/logrus"
	"time"
)

func Info(msg string, data interface{}) {

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.WithFields(logrus.Fields{
		"data": data,
	}).Infoln(msg)
}

func Error(msg string, functionName string, err error, data interface{}) {

	logrus.WithFields(logrus.Fields{
		"level": "error",
		"err":   err,
	}).WithTime(time.Now()).Errorln(msg)
}
func Panic(msg string, functionName string, err error, data interface{}) {

	logrus.WithFields(logrus.Fields{
		"level":        "fatal",
		"Error":        err,
		"FunctionName": functionName,
		"Data":         data,
	}).WithTime(time.Now()).Panicln(msg)
}

func Fatal(msg string, err error) {
	logrus.Fatalln(msg, err)
}
