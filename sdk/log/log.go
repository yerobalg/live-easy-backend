package log

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	log *logrus.Logger
}

func Init() *Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fileInfo(8)
		},
	})
	logger.SetReportCaller(true)

	return &Logger{
		log: logger,
	}
}

func (l *Logger) Debug(ctx context.Context, message string, fields ...interface{}) {
	fmt.Println(fields...)
	l.log.WithContext(ctx).Debug(message)
}

func (l *Logger) Info(ctx context.Context, message string, fields ...interface{}) {
	fmt.Println(fields...)
	l.log.WithContext(ctx).Info(message)
}

func (l *Logger) Warn(ctx context.Context, message string, fields ...interface{}) {
	fmt.Println(fields...)
	l.log.WithContext(ctx).Warn(message)
}

func (l *Logger) Error(ctx context.Context, message string, fields ...interface{}) {
	fmt.Println(fields...)
	l.log.WithContext(ctx).Error(message)
}

func (l *Logger) Fatal(ctx context.Context, message string, fields ...interface{}) {
	fmt.Println(fields...)
	l.log.WithContext(ctx).Fatal(message)
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
	} else {
		location := strings.Split(file, "/")
		file = strings.Join(location[len(location)-3:], "/")
	}
	return fmt.Sprintf("%s:%d", file, line)
}
