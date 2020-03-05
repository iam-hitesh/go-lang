package Routes

import (
	"github.com/gin-gonic/gin"

	"../Controllers"
)


func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/", Controllers.Home)
		v1.GET("/about", Controllers.About)

		v1.POST("/task/create", Controllers.CreateTask)
		v1.GET("/task/list", Controllers.GetAllTasks)
	}

	return router
}

