package routes

import (
	"git.intelbras.com.br/isec/linha-future/jovens/gestoredu.git/api/controller"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AppRoutes(app *fiber.App) {
	usersController := controller.NewUsersController(&gorm.DB{})

	v1 := app.Group("/v1")
	v1.Post("/user", UsersController.CreateUser)
	//v1.Patch("/user/:id", usersController.Patch)
}
