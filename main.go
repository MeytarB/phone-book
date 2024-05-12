package main

import (
	"github.com/MeytarB/phone-book/controller"
	"github.com/MeytarB/phone-book/service/mongo"
	"github.com/gin-gonic/gin"
)


func main() {
	mongoService := mongo.Init()
	phoneBookController := controller.New(mongoService)
	server := gin.Default()
	basepath := server.Group("/app")
	
	phoneBookController.RegisterAllRoutes(basepath)
	server.Run(":3000")
}