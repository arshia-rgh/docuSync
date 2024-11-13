package main

import (
	"context"
	"docuSync/ent"
	"docuSync/logger"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"log"
)

const webPort = 8080

type Config struct {
	client *ent.Client
	logger *logger.Logger
}

func main() {
	// DB and ent initialization
	client := initDB()
	defer client.Close()

	if err := client.Schema.Create(context.TODO()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Println("migration done")

	// Zap logger initialization
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()
	logger_ := logger.New(zapLogger)

	app := Config{client: client, logger: logger_}

	server := fiber.New()
	app.registerRouter(server)

	err := server.Listen(fmt.Sprintf(":%v", webPort))
	if err != nil {
		log.Fatalln(err)
	}
}
