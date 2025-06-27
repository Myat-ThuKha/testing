package main

import (
	"fmt"
	"testing-api/controller"
	"testing-api/database"
	es "testing-api/elasticsearch"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.ConnectMongo()
	if err != nil {
		panic(err)
	}

	es, err := es.NewElasticClient()
	if err != nil {
		panic(err)
	}

	fmt.Println("connected to es :", es)

	r := gin.Default()
	r.POST("/login", controller.LoginHandler(db))
	r.Run()
}
