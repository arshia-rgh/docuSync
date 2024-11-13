package logger

import "go.uber.org/zap"

type Logger struct {
	sugar *zap.SugaredLogger
}

func New(logger *zap.Logger) *Logger {
	return &Logger{sugar: logger.Sugar()}
}

func (l *Logger) Log(level string, msg string, data map[string]any) {
	kvPairs := mapToKV(data)

	switch level {
	case "info":
		l.sugar.Infow(msg, kvPairs...)
	case "warning":
		l.sugar.Warnw(msg, kvPairs...)
	case "error":
		l.sugar.Errorw(msg, kvPairs...)
	default:
		l.sugar.Infow(msg, kvPairs...)

	}
}

func mapToKV(data map[string]any) []any {
	var kvPairs []any
	for k, v := range data {
		kvPairs = append(kvPairs, k, v)
	}

	return kvPairs
}
