package controllers

import (
	"ahiweb/ahiweb/database"
	"ahiweb/ahiweb/models"

	"github.com/gin-gonic/gin"
)

func GetAllPost(c *gin.Context) {
	db := database.GlobalDB
	posts := []models.Posts{}
	db.Find(&posts)
	c.JSON(200, gin.H{
		"Posts": posts,
	})
}

func GetAllUser(c *gin.Context) {
	db := database.GlobalDB
	users := []models.Users{}
	db.Find(&users)
	c.JSON(200, gin.H{
		"user": users,
	})
}
func GetAllF(c *gin.Context) {
	db := database.GlobalDB
	posts := []models.Posts{}
	db.Preload("Users").Find(&posts)

	c.JSON(200, gin.H{
		"Posts": posts,
	})

}
