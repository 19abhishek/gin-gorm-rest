package controller

import (
	"github.com/19abhishek/gin-gorm-rest/config"
	"github.com/19abhishek/gin-gorm-rest/models"
	"github.com/gin-gonic/gin"
)

func UserController(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}