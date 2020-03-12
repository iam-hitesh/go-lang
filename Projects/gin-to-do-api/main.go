package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_"github.com/go-sql-driver/mysql"

	"./Config"
	"./Models"
	"./Routes"
)

var err error

func main() {
	// DB setup with MySQL
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	// If any error found
	if err != nil {
		fmt.Println("Status: ", err)
	}


	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Task{})

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := Routes.SetupRouter()

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	router.Run(":8080")
}