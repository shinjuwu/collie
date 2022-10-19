package log

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	baseLogger *logrus.Logger
	baseFile   *rotatelogs.RotateLogs
}

func New(pathname string) (*Logger, error) {
	// logger
	var baseLogger *logrus.Logger
	var baseFile *rotatelogs.RotateLogs
	if pathname != "" {
		now := time.Now()

		filename := fmt.Sprintf("%d%02d%02d.log",
			now.Year(),
			now.Month(),
			now.Day())
		path := path.Join(filepath.Dir(pathname), filename)
		fileWriter, _ := rotatelogs.New(
			path+".%Y%m%d",
			rotatelogs.WithLinkName(path),
			rotatelogs.WithMaxAge(time.Duration(336)*time.Hour),
			rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
		)
		baseLogger = logrus.New()
		baseFile = fileWriter

	} else {
		baseLogger = logrus.New()
	}
	// new
	logger := new(Logger)
	logger.baseLogger = baseLogger
	logger.baseFile = baseFile
	//Setting
	logger.baseLogger.SetOutput(io.MultiWriter(os.Stdout, logger.baseFile))
	// logger.baseLogger.SetFormatter(&nested.Formatter{
	// 	TimestampFormat: time.RFC3339,
	// })
	return logger, nil
}

// It's dangerous to call the method on logging
func (logger *Logger) Close() {
	if logger.baseFile != nil {
		logger.baseFile.Close()
	}

	logger.baseLogger = nil
	logger.baseFile = nil
}

func (logger *Logger) Debug(format string, a ...interface{}) {
	logger.baseLogger.Debugf(format, a...)
}

func (logger *Logger) Info(format string, a ...interface{}) {
	logger.baseLogger.Infof(format, a...)
}

func (logger *Logger) Release(format string, a ...interface{}) {
	logger.baseLogger.Infof(format, a...)
}

func (logger *Logger) Error(format string, a ...interface{}) {
	logger.baseLogger.Errorf(format, a...)
}

func (logger *Logger) Fatal(format string, a ...interface{}) {
	logger.baseLogger.Fatalf(format, a...)
}

var gLogger, _ = New("")

// It's dangerous to call the method on logging
func Export(logger *Logger) {
	if logger != nil {
		gLogger = logger
	}
}

func Debug(format string, a ...interface{}) {
	gLogger.baseLogger.Debugf(format, a...)
}

func Release(format string, a ...interface{}) {
	gLogger.baseLogger.Infof(format, a...)
}

func Info(format string, a ...interface{}) {
	gLogger.baseLogger.Infof(format, a...)
}

func Error(format string, a ...interface{}) {
	gLogger.baseLogger.Errorf(format, a...)
}

func Fatal(format string, a ...interface{}) {
	gLogger.baseLogger.Fatalf(format, a...)
}

func Close() {
	gLogger.Close()
}
