package main

import (
	"context"
	"docuSync/ent"
	_ "docuSync/ent/runtime"
	"docuSync/logger"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"go.uber.org/fx"
	"log"
)

const webPort = 8080

type Config struct {
	fx.In
	Client *ent.Client
	Logger *logger.Logger
}

func main() {
	app := fx.New(
		AppModule,
		fx.Invoke(registerHooks),
	)
	app.Run()
}

func registerHooks(lc fx.Lifecycle, cfg Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// fiber server initialization
			server := fiber.New()
			server.Get("/metrics", monitor.New())
			server.Use(pprof.New())
			server.Use(cfg.requestLogger)
			server.Use(healthcheck.New())
			server.Use(helmet.New())

			// protected apis ( by auth )
			protectedAPIS := server.Group("/api/protected")
			protectedAPIS.Use(cfg.authenticate)
			cfg.registerProtectedRouter(protectedAPIS)

			// public apis
			publicAPIS := server.Group("/api")
			cfg.registerPublicRouter(publicAPIS)

			go func() {
				err := server.Listen(fmt.Sprintf(":%v", webPort))
				if err != nil {
					log.Fatalln(err)
				}
			}()
			return nil
		},
		OnStop: nil,
	})
}
