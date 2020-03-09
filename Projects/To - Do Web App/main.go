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

	// Environment variable setup
	os.Setenv("PORT", "8080")
	os.Setenv("SECRET_KEY", "PvqcWtGg3OB6wvQSIqqsLfQ9uYV8jMaD")

	// If any error found
	if err != nil {
		fmt.Println("Status: ", err)
	}


	defer Config.DB.Close()
	Config.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Models.User{}, &Models.Task{})
	Config.DB.Model(&Models.Task{}).AddForeignKey("created_by", "user(id)", "RESTRICT", "RESTRICT")

	// Logging to a file.
	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file)

	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	router := Routes.SetupRouter()

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	router.Run(":"+ os.Getenv("PORT"))
}