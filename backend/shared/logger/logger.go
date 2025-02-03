package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"sync"

	"github.com/sirupsen/logrus"
)

// Logger - это структура логгера
type Logger struct {
	*logrus.Logger
}

var (
	instance *Logger
	once     sync.Once
)

func GetLogger() *Logger {
	once.Do(func() {
		log := logrus.New()

		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			ForceColors:     true,
			TimestampFormat: "2006-01-02 15:04:05",
		})

		log.SetLevel(logrus.DebugLevel)
		log.SetOutput(os.Stdout)

		instance = &Logger{log}
	})
	return instance
}

func (l *Logger) logWithCaller(level logrus.Level, msg string) {
	colors := []string{
		logrus.ErrorLevel: "\033[31m", // Красный для ошибок
		logrus.WarnLevel:  "\033[33m", // Желтый для предупреждений
		logrus.InfoLevel:  "\033[34m", // Синий для информационных сообщений
	}

	msg = fmt.Sprintf("%s%s\033[0m", colors[level], msg)
	_, file, line, ok := runtime.Caller(2)

	if !ok {
		l.Logger.Log(level, msg)
	}

	fileLine := fmt.Sprintf("%s:%d", path.Base(file), line)
	coloredFileLine := fmt.Sprintf("\033[32m%s\033[0m", fileLine)
	l.Logger.Log(level, fmt.Sprintf("%s %s", coloredFileLine, msg))
}

func (l *Logger) Error(msg string) { l.logWithCaller(logrus.ErrorLevel, msg) }
func (l *Logger) Info(msg string)  { l.logWithCaller(logrus.InfoLevel, msg) }
func (l *Logger) Warn(msg string)  { l.logWithCaller(logrus.WarnLevel, msg) }
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logWithCaller(logrus.ErrorLevel, fmt.Sprintf(format, args...))
}
func (l *Logger) Infof(format string, args ...interface{}) {
	l.logWithCaller(logrus.InfoLevel, fmt.Sprintf(format, args...))
}
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logWithCaller(logrus.WarnLevel, fmt.Sprintf(format, args...))
}
