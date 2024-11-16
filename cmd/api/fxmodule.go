package main

import (
	"context"
	"docuSync/ent"
	"docuSync/logger"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func provideEntClient() (*ent.Client, error) {
	client := initDB()
	if err := client.Schema.Create(context.TODO()); err != nil {
		return nil, err
	}
	return client, nil
}

func provideLogger() (*logger.Logger, error) {
	zapLogger := zap.NewExample()
	logger_ := logger.New(zapLogger)
	return logger_, nil
}

var AppModule = fx.Options(
	fx.Provide(provideEntClient),
	fx.Provide(provideLogger),
)
