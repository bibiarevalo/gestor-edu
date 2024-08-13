package database

import (
	"fmt"

	"git.intelbras.com.br/isec/linha-future/jovens/gestoredu.git/api/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Db() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/defense"
	//var v = "Não conseguiu conectar ao banco de dados"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB = db
	db.AutoMigrate(&entities.Users{})
	db.AutoMigrate(&entities.Administrador{})
	db.AutoMigrate(&entities.Aluno{})
	db.AutoMigrate(&entities.Professor{})
	db.AutoMigrate(&entities.Responsavel{})
	// if err != nil {
	// 	panic(v)
	// }
	fmt.Println("conexão OK!")
	fmt.Println(db)

}
