package main

import (
	"testing-api/controller"
	"testing-api/database"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.ConnectMongo()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.POST("/login", controller.LoginHandler(db))
	r.Run()
}
