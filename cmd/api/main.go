package main

import (
	"context"
	"docuSync/ent"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

const webPort = 8080

type Config struct {
	client *ent.Client
}

func main() {
	client := initDB()
	defer client.Close()

	if err := client.Schema.Create(context.TODO()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Println("migration done")
	server := fiber.New()

	err := server.Listen(fmt.Sprintf(":%v", webPort))
	if err != nil {
		log.Fatalln(err)
	}
}
