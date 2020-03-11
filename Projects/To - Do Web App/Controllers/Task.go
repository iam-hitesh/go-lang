package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"../Models"
	"../Services"
)


func CreateTask(req *gin.Context) {
	var task Models.Task

	// Should be with conditional check otherwise give error:
	// Headers were already written. Wanted to override status code 400 with 200
	// Content-Type: application/json
	if err := req.Bind(&task); err != nil {
		return
	}

	currentUser := (req.MustGet("currentUser")).(map[string]interface{})

	task.CreatedBy = uint(currentUser["ID"].(float64))

	err := task.Create()

	if err != nil {
		req.AbortWithStatus(http.StatusBadRequest)
	} else {
		req.JSON(http.StatusCreated, task)
	}
}


func GetATask(req *gin.Context) {
	var task Models.Task
	currentUser := (req.MustGet("currentUser")).(map[string]interface{})

	id := req.Params.ByName("id")

	err := task.Get(id)

	if task.CreatedBy != uint(currentUser["ID"].(float64)){
		req.Abort()
		req.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not accessed by other than created by",
		})

		return
	}

	if err != nil {
		req.AbortWithStatus(http.StatusNotFound)
	} else {
		req.JSON(http.StatusOK, task)
	}
}


func UpdateATask(req *gin.Context) {
	var task Models.Task
	currentUser := (req.MustGet("currentUser")).(map[string]interface{})

	id := req.Params.ByName("id")

	err := task.Get(id)

	if err != nil {
		req.Abort()
		req.JSON(http.StatusNotFound, gin.H{
			"message": "Task not found",
		})

		return
	}

	if task.CreatedBy != uint(currentUser["ID"].(float64)){
		req.Abort()
		req.JSON(http.StatusUnauthorized, gin.H{
			"message": "You are not authorized to update this task",
		})

		return
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
	currentUser := (req.MustGet("currentUser")).(map[string]interface{})

	id := req.Params.ByName("id")

	err := task.Delete(id, uint(currentUser["ID"].(float64)))

	if err != nil {
		req.AbortWithStatus(http.StatusNotFound)
	} else {
		req.JSON(http.StatusOK, task)
	}
}


func GetAllTasks(req *gin.Context) {
	var tasks []Models.Task
	currentUser := (req.MustGet("currentUser")).(map[string]interface{})

	err := Services.GetAllTasks(&tasks, uint(currentUser["ID"].(float64)))

	if err != nil {
		req.AbortWithStatus(http.StatusBadRequest)
	} else {
		req.JSON(http.StatusOK, tasks)
	}
}