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

// GetLogger возвращает единственный экземпляр логгера
func GetLogger() *Logger {
	once.Do(func() {
		log := logrus.New()

		// Устанавливаем формат вывода
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			ForceColors:     true,
			TimestampFormat: "2006-01-02 15:04:05", // Изменяем формат времени
		})

		// Устанавливаем уровень логирования
		log.SetLevel(logrus.DebugLevel)

		// Устанавливаем вывод в консоль
		log.SetOutput(os.Stdout)

		instance = &Logger{log} // Инициализируем экземпляр логгера
	})
	return instance
}

// logWithCaller добавляет информацию о вызове
func (l *Logger) logWithCaller(level logrus.Level, msg string) {
	var color string
	switch level {
	case logrus.ErrorLevel:
		color = "\033[31m" // Красный для ошибок
	case logrus.WarnLevel:
		color = "\033[33m" // Желтый для предупреждений
	case logrus.InfoLevel:
		color = "\033[34m" // Синий для информационных сообщений
	}

	msg = fmt.Sprintf("%s%s\033[0m", color, msg)
	_, file, line, ok := runtime.Caller(2) // Используем 2, чтобы получить информацию о вызове логгера
	if ok {
		fileLine := fmt.Sprintf("%s:%d", path.Base(file), line)

		// Добавляем ANSI-коды для цвета (например, зеленый)
		coloredFileLine := fmt.Sprintf("\033[32m%s\033[0m", fileLine) // Зеленый цвет

		l.Logger.Log(level, fmt.Sprintf("%s %s", coloredFileLine, msg))
	} else {
		l.Logger.Log(level, msg)
	}
}

// Error добавляет информацию о месте вызова и логирует сообщение об ошибке
func (l *Logger) Error(msg string) {
	l.logWithCaller(logrus.ErrorLevel, msg)
}

// Info добавляет информацию о месте вызова и логирует информационное сообщение
func (l *Logger) Info(msg string) {
	l.logWithCaller(logrus.InfoLevel, msg)
}

// Warn добавляет информацию о месте вызова и логирует предупреждение
func (l *Logger) Warn(msg string) {
	l.logWithCaller(logrus.WarnLevel, msg)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logWithCaller(logrus.ErrorLevel, fmt.Sprintf(format, args...))
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.logWithCaller(logrus.InfoLevel, fmt.Sprintf(format, args...))
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logWithCaller(logrus.WarnLevel, fmt.Sprintf(format, args...))
}
