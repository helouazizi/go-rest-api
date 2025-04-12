package dependencies

import (
	"go-rest-api/internal/handlers"
	"go-rest-api/internal/services"
)

type Dependencies struct {
	// Services
	ItemService *services.ItemService
	// UserService *services.UserService

	// Handlers
	ItemHandler *handlers.ItemHandler
	// UserHandler *handlers.UserHandler
}

func NewDependencies() *Dependencies {
	// Instantiate services
	itemService := services.NewItemService()
	// userService := services.NewUserService()

	// Instantiate handlers
	itemHandler := handlers.NewItemHandler(itemService)
	// userHandler := handlers.NewUserHandler(userService)

	return &Dependencies{
		ItemService: itemService,
		// UserService:  userService,
		ItemHandler: itemHandler,
		// UserHandler:  userHandler,
	}
}
