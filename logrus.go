package logger

import (
	"github.com/sirupsen/logrus"
)

type loggersImpl struct {
	impl logrus.FieldLogger
}

func New() Logger {
	log := logrus.New()
	return NewWithLogger(log)
}

func NewWithLogger(logger *logrus.Logger) Logger {
	return &loggersImpl{impl: logrus.NewEntry(logger)}
}

func NewWithEntry(logger *logrus.Entry) Logger {
	return &loggersImpl{impl: logger}
}

func NewWithFieldLogger(logger logrus.FieldLogger) Logger {
	return &loggersImpl{impl: logger}
}

func (l *loggersImpl) Fatal(format ...interface{}) {
	l.impl.Fatal(format...)
}

func (l *loggersImpl) Fatalf(format string, args ...interface{}) {
	l.impl.Fatalf(format, args...)
}

func (l *loggersImpl) Fatalln(args ...interface{}) {
	l.impl.Fatalln(args...)
}

func (l *loggersImpl) Debug(args ...interface{}) {
	l.impl.Debug(args...)
}

func (l *loggersImpl) Debugf(format string, args ...interface{}) {
	l.impl.Debugf(format, args...)
}

func (l *loggersImpl) Debugln(args ...interface{}) {
	l.impl.Debugln(args...)
}

func (l *loggersImpl) Error(args ...interface{}) {
	l.impl.Error(args...)
}

func (l *loggersImpl) Errorf(format string, args ...interface{}) {
	l.impl.Errorf(format, args...)
}

func (l *loggersImpl) Errorln(args ...interface{}) {
	l.impl.Errorln(args...)
}

func (l *loggersImpl) Info(args ...interface{}) {
	l.impl.Info(args...)
}

func (l *loggersImpl) Infof(format string, args ...interface{}) {
	l.impl.Infof(format, args...)
}

func (l *loggersImpl) Infoln(args ...interface{}) {
	l.impl.Infoln(args...)
}

func (l *loggersImpl) Warn(args ...interface{}) {
	l.impl.Warn(args...)
}

func (l *loggersImpl) Warnf(format string, args ...interface{}) {
	l.impl.Warnf(format, args...)
}

func (l *loggersImpl) Warnln(args ...interface{}) {
	l.impl.Warnln(args...)
}

func (l *loggersImpl) Panic(args ...interface{}) {
	l.impl.Panic(args...)
}

func (l *loggersImpl) Panicf(format string, args ...interface{}) {
	l.impl.Panicf(format, args...)
}

func (l *loggersImpl) Panicln(args ...interface{}) {
	l.impl.Panicln(args...)
}

func (l *loggersImpl) WithFields(fields map[string]interface{}) Logger {
	return &loggersImpl{impl: l.impl.WithFields(fields)}
}

func (l *loggersImpl) WithField(key string, value interface{}) Logger {
	return &loggersImpl{impl: l.impl.WithField(key, value)}
}

func (l *loggersImpl) WithError(err error) Logger {
	return &loggersImpl{impl: l.impl.WithError(err)}
}
