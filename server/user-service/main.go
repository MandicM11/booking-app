package main

import (
	"log"

	"booking-app/server/user-service/handler"
	"booking-app/server/user-service/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	repository.Connect()
	go routerSetup()
	//grpcserver.InitServer()
	repository.Disconnect()
}

func routerSetup() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.NoRoute()

	routes := router.Group("/user")
	routes.GET("/all", handler.GetAll)
	routes.GET("/id/:id", handler.GetById)
	routes.GET("/email/:email", handler.GetByEmail)
	routes.POST("/register/admin", handler.SignupAdmin)
	routes.POST("/register/guest", handler.SignupGuest)
	routes.POST("/update", handler.UpdateUser)
	routes.POST("/login", handler.Login)

	//routes.DELETE("/delete/:id", handler.DeleteUser)
	router.Run(":8080")
	log.Println("HTTP server running on port 8080")
}
