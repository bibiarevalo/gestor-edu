package models

type User struct {
	ID              string
	Nome            string
	Sobrenome       string
	Telefone        string
	Email           string
	CPF             string
	Endereço        string
	Idade           int
	Data_nascimento string
	Matricula       string
	Tipo            string
}

//pertence ao package controller, create 
//user.ID = uuid.New().String()
//abaixo dessa função

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