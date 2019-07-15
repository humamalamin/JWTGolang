package config

import (
	"JWTGolang/structs"
	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/jwt-golang")
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(structs.Person{})
	return db
}
