package controller

import (
	"net/http"
	"git.intelbras.com.br/isec/linha-future/jovens/gestoredu.git/api/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UsersController struct {
	users []interface{}
	db    *gorm.DB
}

func NewUsersController(db *gorm.DB) *UsersController {
	return &UsersController{db: db}
}

func (uc *UsersController) CreateUser(c *fiber.Ctx) error {
	var user entities.Users

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse body",
		})
	}

	user.ID = uuid.New().String()

	// var response interface{}

	// switch user.Tipo {
	// case "Administrador":
	// 	var admin entities.Administrador
	// 	admin.User = user

	// 	if err := c.BodyParser(&admin); err != nil {
	// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
	// 			"error": "Failed to parse additional fields for Administrador",
	// 		})
	// 	}
	// 	response = admin

	// case "Professor":
	// 	var professor entities.Professor
	// 	professor.User = user

	// 	if err := c.BodyParser(&professor); err != nil {
	// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
	// 			"error": "Failed to parse additional fields for Professor",
	// 		})
	// 	}
	// 	response = professor

	// case "Aluno":
	// 	var aluno entities.Aluno
	// 	aluno.User = user

	// 	if err := c.BodyParser(&aluno); err != nil {
	// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
	// 			"error": "Failed to parse additional fields for Aluno",
	// 		})
	// 	}
	// 	uc.db.Create(&aluno)
	// 	response = aluno

	// case "Responsavel":
	// 	var responsavel entities.Responsavel
	// 	responsavel.User = user

	// 	if err := c.BodyParser(&responsavel); err != nil {
	// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
	// 			"error": "Failed to parse additional fields for Responsavel",
	// 		})
	// 	}
	// 	response = responsavel

	// default:
	// 	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
	// 		"error": "Invalid user type",
	// 	})
	// }
	uc.db.Save(&user)
	return c.Status(http.StatusCreated).JSON(&user)
}
