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
	// Content-Type: application/json
	if err := req.Bind(&task); err != nil {
		return
	}

	err := task.Create()

	if err != nil {
		req.AbortWithStatus(http.StatusBadRequest)
	} else {
		req.JSON(http.StatusCreated, task)
	}
}


func GetATask(req *gin.Context) {
	var task Models.Task

	id := req.Params.ByName("id")

	err := task.Get(id)

	if err != nil {
		req.AbortWithStatus(http.StatusNotFound)
	} else {
		req.JSON(http.StatusOK, task)
	}
}


func UpdateATask(req *gin.Context) {
	var task Models.Task

	id := req.Params.ByName("id")

	err := task.Get(id)

	if err != nil {
		req.JSON(http.StatusNotFound, task)
	}

	req.BindJSON(&task)
	err = task.Update()

	if err != nil {
		req.AbortWithStatus(http.StatusBadRequest)
	} else {
		req.JSON(http.StatusOK, task)
	}
}


func DeleteATask(req *gin.Context) {
	var task Models.Task

	id := req.Params.ByName("id")

	err := task.Delete(id)

	if err != nil {
		req.AbortWithStatus(http.StatusNotFound)
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