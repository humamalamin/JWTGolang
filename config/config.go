package config

import (
	"JWTGolang/structs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=jwt-golang password=root")
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(structs.Person{})
	return db
}
