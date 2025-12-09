package router

import (
	"ahiweb/ahiweb/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/test", controllers.GetAllUser)
	r.GET("/test1", controllers.GetAllPost)
	r.GET("/a", controllers.GetAllF)
	return r
}
