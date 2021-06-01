package cmd

import (
	"context"
	"customer/api/controller"
	"customer/domain"
	"customer/domain/model"
	"customer/internal"
	"customer/repository"
	"customer/service"
	"log"
)

// initDatabase initialize the database repository
func initDatabase() internal.Database {
	db, err := internal.NewMemDatabaseImpl()
	if err != nil {
		panic("Error when creating new database instances")
	}

	// create migration table
	if err := db.CreateTable(model.Customer{}); err != nil {
		log.Println("Error when creating the table")
	}

	return db
}

// initRepo initialize the repository
func initRepo(db internal.Database) domain.CustomerRepository {
	return repository.CustomerRepositoryImpl{
		Ctx: context.Background(),
		DB:  db,
	}
}

func initService(repo domain.CustomerRepository) domain.Services {
	return domain.Services{
		AuthService:     service.AuthServiceImpl{Repository: repo},
		CustomerService: service.CustomerServiceImpl{Repository: repo},
	}
}

// initEntities initialize the database entities
func initEntities(db internal.Database) *controller.Controller {
	customerRepo := initRepo(db)
	customerService := initService(customerRepo)

	customerController := &controller.CustomerController{Service: customerService.CustomerService}
	authController := &controller.AuthController{Service: customerService.AuthService}

	return &controller.Controller{
		CustomerController: customerController,
		AuthController:     authController,
	}
}
