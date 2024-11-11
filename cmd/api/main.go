package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

const webPort = 8080

type Config struct{}

func main() {
	server := fiber.New()

	err := server.Listen(fmt.Sprintf(":%v", webPort))
	if err != nil {
		log.Fatalln(err)
	}
}
