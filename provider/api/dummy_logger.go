package api

import (
	"io"

	"github.com/labstack/gommon/log"
)

type dummyLogger struct{}

func (d dummyLogger) Output() io.Writer                         { return dummyWriter{} }
func (d dummyLogger) SetOutput(w io.Writer)                     {}
func (d dummyLogger) Prefix() string                            { return "" }
func (d dummyLogger) SetPrefix(p string)                        {}
func (d dummyLogger) Level() log.Lvl                            { return log.OFF }
func (d dummyLogger) SetLevel(v log.Lvl)                        {}
func (d dummyLogger) SetHeader(h string)                        {}
func (d dummyLogger) Print(i ...interface{})                    {}
func (d dummyLogger) Printf(format string, args ...interface{}) {}
func (d dummyLogger) Printj(j log.JSON)                         {}
func (d dummyLogger) Debug(i ...interface{})                    {}
func (d dummyLogger) Debugf(format string, args ...interface{}) {}
func (d dummyLogger) Debugj(j log.JSON)                         {}
func (d dummyLogger) Info(i ...interface{})                     {}
func (d dummyLogger) Infof(format string, args ...interface{})  {}
func (d dummyLogger) Infoj(j log.JSON)                          {}
func (d dummyLogger) Warn(i ...interface{})                     {}
func (d dummyLogger) Warnf(format string, args ...interface{})  {}
func (d dummyLogger) Warnj(j log.JSON)                          {}
func (d dummyLogger) Error(i ...interface{})                    {}
func (d dummyLogger) Errorf(format string, args ...interface{}) {}
func (d dummyLogger) Errorj(j log.JSON)                         {}
func (d dummyLogger) Fatal(i ...interface{})                    {}
func (d dummyLogger) Fatalj(j log.JSON)                         {}
func (d dummyLogger) Fatalf(format string, args ...interface{}) {}
func (d dummyLogger) Panic(i ...interface{})                    {}
func (d dummyLogger) Panicj(j log.JSON)                         {}
func (d dummyLogger) Panicf(format string, args ...interface{}) {}

type dummyWriter struct{}

func (d dummyWriter) Write(p []byte) (n int, err error) { return 0, nil }
