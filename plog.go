package plog

import (
	"context"
	"encoding/json"
	"io"
	"strings"
	"time"

	"github.com/go-puzzles/plog/level"
	"github.com/go-puzzles/plog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger Logger
)

func init() {
	time.Local = time.FixedZone("CST", 8*3600)
	logger = log.New(log.WithCalldepth(4))
}

func SetLogger(l Logger) {
	logger = l
}

func IsDebug() bool {
	return logger.IsDebug()
}

func EnableLogToFile(jackLog *LogConfig) {
	logger.SetOutput((*lumberjack.Logger)(jackLog))
}

func SetOutput(w io.Writer) {
	logger.SetOutput(w)
}

func Enable(l level.Level) {
	logger.Enable(l)
}

func Errorf(msg string, v ...any) {
	logger.Errorf(msg, v...)
}

func Warnf(msg string, v ...any) {
	logger.Warnf(msg, v...)
}

func Infof(msg string, v ...any) {
	logger.Infof(msg, v...)
}

func Debugf(msg string, v ...any) {
	logger.Debugf(msg, v...)
}

func Fatalf(msg string, v ...any) {
	logger.Fatalf(msg, v...)
}

func Infoc(ctx context.Context, msg string, v ...any) {
	logger.Infoc(ctx, msg, v...)
}

func Debugc(ctx context.Context, msg string, v ...any) {
	logger.Debugc(ctx, msg, v...)
}

func Warnc(ctx context.Context, msg string, v ...any) {
	logger.Warnc(ctx, msg, v...)
}

func Errorc(ctx context.Context, msg string, v ...any) {
	logger.Errorc(ctx, msg, v...)
}

// TimeFuncDuration returns the duration consumed by function.
// It has specified usage like:
//
//	    f := TimeFuncDuration()
//		   DoSomething()
//		   duration := f()
func TimeFuncDuration() func() time.Duration {
	start := time.Now()
	return func() time.Duration {
		return time.Since(start)
	}
}

func TimeDurationDefer(prefix ...string) func() {
	ps := "operation"
	if len(prefix) != 0 {
		ps = strings.Join(prefix, ", ")
	}
	start := time.Now()

	return func() {
		Infof("%v elapsed time: %v", ps, time.Since(start))
	}
}

func Jsonify(v any) string {
	d, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		logger.Errorf("jsonify error: %v", err)
		panic(err)
	}
	return string(d)
}
