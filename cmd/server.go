package cmd

import (
	"customer/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Serve() {
	app := fiber.New()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c

		_ = <-c
		_ = app.Shutdown()

		os.Exit(0)
	}()

	db := initDatabase()
	defer func() {
		if err := db.Close(); err != nil {
			log.Println("Error when close db connection ", err)
		}
	}()

	entities := initEntities(db)

	app.Use(recover.New())
	app.Use(logger.New())

	api.RegisterRoute(app, entities)

	if err := app.Listen(":3000"); err != nil {
		log.Println("Error when listen from the HTTP Framework ", err)
	}
}
