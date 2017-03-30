package main

import (
	"os"

	"github.com/Sirupsen/logrus"
)

type logger struct {
	logHandler *logrus.Logger
	logEnable  bool
}

var Log = &logger{}

func (l *logger) Init(logEnable bool) {
	l.logEnable = logEnable
	l.logHandler = logrus.New()
	formatter := new(logrus.TextFormatter)
	formatter.FullTimestamp = true
	l.logHandler.Formatter = formatter // default
}

//there are two type of errors:
//FATAL (exit program)
//WARNING (can continue)
func (l *logger) WriteLine(typeErr, action string, err error) {
	if !l.logEnable {
		return
	}
	if typeErr == "LOG" {
		l.logHandler.WithFields(logrus.Fields{}).Info(action)
		return
	}
	if err == nil {
		return
	}
	if typeErr == "WARNING" {
		l.logHandler.WithFields(logrus.Fields{}).Warn(action + " -> " + err.Error())
		return
	}
	l.logHandler.WithFields(logrus.Fields{}).Error(action + " -> " + err.Error())
	os.Exit(1)
}
