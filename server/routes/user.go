package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nickel-dime/user-management/controllers"
)

// sets the routes for the server
func UserManagementRoutes(route fiber.Router) {
	route.Get("", controllers.GetAllUsers)
	route.Post("", controllers.CreateUser)
	route.Post("/:id", controllers.UpdateUser)
	route.Delete("/:id", controllers.DeleteUser)
}
