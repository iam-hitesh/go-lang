package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"../Models"
)


func CreateTask(req *gin.Context) {
	var task Models.Task

	// Should be with conditional check otherwise give error:
	// Headers were already written. Wanted to override status code 400 with 200
	if err := req.Bind(&task); err != nil {
		return
	}

	err := Models.CreateTask(&task)

	if err != nil {
		req.AbortWithStatus(http.StatusBadRequest)
	} else {
		req.JSON(http.StatusOK, task)
	}
}


func GetAllTasks(req *gin.Context) {
	var tasks []Models.Task

	err := Models.GetAllTasks(&tasks)

	if err != nil {
		req.AbortWithStatus(http.StatusBadRequest)
	} else {
		req.JSON(http.StatusOK, tasks)
	}
}