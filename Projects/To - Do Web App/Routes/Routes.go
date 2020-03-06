package Routes

import (
	"github.com/gin-gonic/gin"

	"../Controllers"
)


func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("", Controllers.Home)
		v1.GET("about", Controllers.About)

		v1.POST("user/create", Controllers.CreateUser)
		v1.GET("user", Controllers.GetAllUsers)

		v1.POST("task/create", Controllers.CreateTask)
		v1.GET("task", Controllers.GetAllTasks)
		v1.GET("task/:id", Controllers.GetATask)
		v1.PUT("task/:id", Controllers.UpdateATask)
		v1.DELETE("task/:id", Controllers.DeleteATask)
	}

	return router
}

