package main

import (
	"git.intelbras.com.br/isec/linha-future/jovens/gestoredu.git/api/routes"
	"git.intelbras.com.br/isec/linha-future/jovens/gestoredu.git/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db := database.Db()
	app := fiber.New()

	routes.AppRoutes(app, db)

	app.Listen(":3000")
}
