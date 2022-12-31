package log

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
	"live-easy-backend/sdk/appcontext"
	"live-easy-backend/sdk/auth"
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

func (l *Logger) Debug(ctx context.Context, message string, field ...interface{}) {
	l.log.WithContext(ctx).WithFields(getFields(ctx, field...)).Debug(message)
}

func (l *Logger) Info(ctx context.Context, message string, field ...interface{}) {
	l.log.WithContext(ctx).WithFields(getFields(ctx, field...)).Info(message)
}

func (l *Logger) Warn(ctx context.Context, message string, field ...interface{}) {
	l.log.WithContext(ctx).WithFields(getFields(ctx, field...)).Warn(message)
}

func (l *Logger) Error(ctx context.Context, message string, field ...interface{}) {
	l.log.WithContext(ctx).WithFields(getFields(ctx, field...)).Error(message)
}

func (l *Logger) Fatal(ctx context.Context, message string, field ...interface{}) {
	l.log.WithContext(ctx).WithFields(getFields(ctx, field...)).Fatal(message)
}

func getFields(ctx context.Context, fields ...interface{}) logrus.Fields {
	logFields := logrus.Fields{
		"request_id":      appcontext.GetRequestId(ctx),
		"service_version": appcontext.GetServiceVersion(ctx),
		"user_agent":      appcontext.GetUserAgent(ctx),
		"user_id":         auth.GetUserID(ctx),
	}

	if len(fields) > 0 {
		logFields["data"] = fields[0]
	} else {
		logFields["data"] = nil
	}

	return logFields
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
