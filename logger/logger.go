package logger

import "go.uber.org/zap"

type Log struct {
	log *zap.Logger
}

func NewLog() *Log {

	log := zap.Must(zap.NewDevelopment())
	defer log.Sync()
	return &Log{
		log: log,
	}
}

func (l *Log) Errorf(msg string, err error) {
	if err != nil {
		l.log.Sugar().Errorf(msg, err)
	}
}

func (l *Log) Error(err error) {
	if err != nil {
		l.log.Sugar().Error(err)
	}
}

func (l *Log) Infof(msg string, err error) {
	if err != nil {
		l.log.Sugar().Infof(msg, err)
	}
}

func (l *Log) Info(err error) {
	if err != nil {
		l.log.Sugar().Info(err)
	}
}

func (l *Log) Debugf(msg string, err error) {
	if err != nil {
		l.log.Sugar().Debugf(msg, err)
	}
}

func (l *Log) Debug(err error) {
	if err != nil {
		l.log.Sugar().Debug(err)
	}
}
