package entities

import (
	"github.com/pborman/uuid"
)

type Users struct {
	ID             string `json:"id"`
	Nome           string `json:"nome"`
	Sobrenome      string `json:"sobrenome"`
	Telefone       int    `json:"telefone"`
	Email          string `json:"email"`
	Sexo           string `json:"sexo"`
	CPF            int    `json:"cpf"`
	Endereco       string `json:"endereco"`
	DataNascimento string `json:"dataNascimento"`
	Matricula      string `json:"matricula"`
	Tipo           string `json:"tipo"`
}

type Administrador struct {
	User      Users `gorm:"foreignKey:UserID;references:ID"`
	UserID    string
	Permissao bool `json:"permissao"`
}

type Professor struct {
	User          Users `gorm:"foreignKey:UserID;references:ID"`
	UserID        string
	Disciplina    string `json:"disciplina"`
	Especialidade string `json:"especialidade"`
	Turno         string `json:"turno"`
}

type Aluno struct {
	User   Users `gorm:"foreignKey:UserID;references:ID"`
	UserID string
	Serie  int `json:"serie"`
}

type Responsavel struct {
	User   Users `gorm:"foreignKey:UserID;references:ID"`
	UserID string
	Aluno  string `json:"aluno"`
}

func NewUser(tipo string) interface{} {
	baseUser := Users{
		ID:   uuid.New(),
		Tipo: tipo,
	}

	switch tipo {
	case "administrador":
		return &Administrador{
			User: baseUser,
		}
	case "professor":
		return &Professor{
			User: baseUser,
		}
	case "aluno":
		return &Aluno{
			User: baseUser,
		}
	case "responsavel":
		return &Responsavel{
			User: baseUser,
		}
	default:
		return nil
	}
}
