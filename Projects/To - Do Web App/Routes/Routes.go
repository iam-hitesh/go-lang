package Routes

import (
	"github.com/gin-gonic/gin"
	"../Controllers"
)

func setupRouter() *gin.Engine{
	routes := gin.Default()

	v1 := routes.Group("/v1")
	{
		v1.GET("task/create", Controllers.createTask)
	}

	return v1
}
