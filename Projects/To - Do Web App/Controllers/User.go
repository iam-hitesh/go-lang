package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"../Models"
	"../Services"
)


func CreateUser(req *gin.Context) {
	var user Models.User

	if err := req.Bind(&user); err != nil {
		return
	}

	err := user.Create()

	if err != nil {
		req.AbortWithStatus(http.StatusBadRequest)
	} else {
		req.JSON(http.StatusCreated, user)
	}
}


func GetAllUsers(req *gin.Context) {
	var users []Models.User

	err := Services.GetAllUsers(&users)

	if err != nil {
		req.AbortWithStatus(http.StatusBadRequest)
	} else {
		req.JSON(http.StatusOK, users)
	}
}


func GetUserProfile(req *gin.Context) {
	user := req.MustGet("currentUser")
	currentUser := user.(map[string]interface{})

	var User Models.User
	User.GetByEmail(currentUser["Email"].(string))

	User.Password = ""

	req.JSON(http.StatusOK, gin.H{
		"user": User,
	})
}