package main

import (
	"JWTGolang/config"
	"JWTGolang/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB:db}

	router := gin.Default()

	router.GET("/person/:id", inDB.GetPerson)
	router.GET("/person", inDB.GetPersons)
	router.POST("/person/create", inDB.CreatePerson)
	router.PUT("/person", inDB.UdatePerson)
	router.DELETE("/person/:id", inDB.DeletePerson)
	router.Run(":3000")
}
