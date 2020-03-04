package Models

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"../Config"
)

type Task struct {
	ID uint 			`json:"id"`
	Title string 		`json:"Title"`
	Description string 	`json:"Description"`
	Completed boolean	`json:"Completed"`
}

func (table *Task) TableName() string {
	return "task"
}

func getAllTasks(tasks *[]Task) (err error) {
	if err = Config.DB.Find(task).Error; err != nil {
		return err
	}

	return nil
}

func createTask(task *Task) (err error) {
	if err = Config.DB.Create(task).Error; err != nil {
		return err
	}

	return nil
}
