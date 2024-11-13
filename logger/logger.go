package logger

import "go.uber.org/zap"

type Logger struct {
	sugar *zap.SugaredLogger
}

func New(logger *zap.Logger) *Logger {
	return &Logger{sugar: logger.Sugar()}
}

func (l *Logger) Log(level string, msg string, data map[string]any) {
	switch level {
	case "info":
		l.infoLog(msg, data)
	case "warning":
		l.warnLog(msg, data)
	case "error":
		l.errorLog(msg, data)
	default:
		l.infoLog(msg, data)

	}
}

func (l *Logger) infoLog(msg string, data map[string]any) {
	l.sugar.Infow(msg, data)
}

func (l *Logger) warnLog(msg string, data map[string]any) {
	l.sugar.Warnw(msg, data)
}

func (l *Logger) errorLog(msg string, data map[string]any) {
	l.sugar.Errorw(msg, data)
}
