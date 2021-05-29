package cmd

import (
	"context"
	"customer/api/controller"
	"customer/domain/model"
	"customer/repository"
	"customer/service"
	"log"
)

// initDatabase initialize the database repository
func initDatabase() repository.Database {
	db, err := repository.NewMemDatabaseImpl()
	if err != nil {
		panic("Error when creating new database instances")
	}

	// create migration table
	if err := db.CreateTable(model.Customer{}); err != nil {
		log.Println("Error when creating the table")
	}

	return db
}

// initEntities initialize the database entities
func initEntities(db repository.Database) *controller.Controller {
	customerRepo := repository.CustomerRepository{
		Ctx: context.Background(),
		DB:  db,
	}
	customerService := service.CustomerService{Repository: customerRepo}
	customerController := &controller.CustomerController{Service: customerService}

	return &controller.Controller{
		CustomerController: customerController,
	}
}
