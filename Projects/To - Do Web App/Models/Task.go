package Models

import (
	"../Config"
)

type Task struct {
	ID uint 			`json:"id"`
	Title string 		`json:"Title"`
	Description string 	`json:"Description"`
	Completed bool	`json:"Completed"`
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

func createTask(t *Task) (err error) {
	if err = Config.DB.Create(task).Error; err != nil {
		return err
	}

	return nil
}
