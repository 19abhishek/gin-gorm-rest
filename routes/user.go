package routes

import (
	"github.com/19abhishek/gin-gorm-rest/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine){
	router.GET("/", controller.UserController)

}