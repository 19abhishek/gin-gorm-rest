package main

import (
	"github.com/19abhishek/gin-gorm-rest/config"
	"github.com/19abhishek/gin-gorm-rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8080");
}