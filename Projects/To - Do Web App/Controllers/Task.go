package Controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"../Models"
)

func createTask(body *gin.context){
	var task Models.Task

	body.BindJSON(&task)

	err := Models.CreateATodo(&todo)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, task)
	}
}