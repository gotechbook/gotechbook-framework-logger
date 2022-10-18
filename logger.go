package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log = initLog()

func initLog() Logger {
	s := &logrus.Logger{
		Out:          os.Stderr,
		Formatter:    new(logrus.TextFormatter),
		Hooks:        make(logrus.LevelHooks),
		Level:        logrus.DebugLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}
	return NewWithFieldLogger(s.WithFields(logrus.Fields{"source": PREFIX}))
}

func Set(l Logger) {
	if l != nil {
		Log = l
	}
}
