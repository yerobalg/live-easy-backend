package log

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm/logger"
)

// TODO: move to new package

// ErrRecordNotFound record not found error
var ErrRecordNotFound = errors.New("record not found")

// LogLevel log level
type LogLevel int

const (
	// Silent silent log level
	Silent LogLevel = iota + 1
	// Error error log level
	Error
	// Warn warn log level
	Warn
	// Info info log level
	Info
)

// Config logger config
type Config struct {
	SlowThreshold             time.Duration
	Colorful                  bool
	IgnoreRecordNotFoundError bool
	ParameterizedQueries      bool
	LogLevel                  LogLevel
}

// Interface logger interface
type Interface interface {
	LogMode(logger.LogLevel) logger.Interface
	Info(context.Context, string, ...interface{})
	Warn(context.Context, string, ...interface{})
	Error(context.Context, string, ...interface{})
	Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error)
}

// New initialize logger
func New(config Config, serverLogger *Logger) Interface {
	return &GormLogger{
		Config:       config,
		ServerLogger: serverLogger,
	}
}

type GormLogger struct {
	Config
	ServerLogger *Logger
}

// LogMode log mode
func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

// Info print info
func (l GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= Info {
		l.ServerLogger.Info(ctx, msg, data...)
	}
}

// Warn print warn messages
func (l GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= Warn {
		l.ServerLogger.Warn(ctx, msg, data...)
	}
}

// Error print error messages
func (l GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= Error {
		l.ServerLogger.Error(ctx, msg, data...)
	}
}

// Trace print sql message
func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= Error && (!errors.Is(err, ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			l.Error(ctx, err.Error(), "elapsed", float64(elapsed.Nanoseconds())/1e6, "rows", "-", "sql", sql)
		} else {
			l.Error(ctx, err.Error(), "elapsed", float64(elapsed.Nanoseconds())/1e6, "rows", rows, "sql", sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			l.Warn(ctx, slowLog, "elapsed", float64(elapsed.Nanoseconds())/1e6, "rows", "-", "sql", sql)
		} else {
			l.Warn(ctx, slowLog, "elapsed", float64(elapsed.Nanoseconds())/1e6, "rows", rows, "sql", sql)
		}
	case l.LogLevel == Info:
		sql, rows := fc()
		if rows == -1 {
			l.Info(ctx, "sql", "elapsed", float64(elapsed.Nanoseconds())/1e6, "rows", "-", "query", sql)
		} else {
			l.Info(ctx, "sql", "elapsed", float64(elapsed.Nanoseconds())/1e6, "rows", rows, "query", sql)
		}
	}
}

// Trace print sql message
func (l GormLogger) ParamsFilter(ctx context.Context, sql string, params ...interface{}) (string, []interface{}) {
	if l.Config.ParameterizedQueries {
		return sql, nil
	}
	return sql, params
}
