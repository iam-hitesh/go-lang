package Models

import (
	"../Config"
)


type Task struct {
	ID uint 			`json:"id"`
	Title string 		`json:"Title"`
	Description string 	`json:"Description"`
	Completed bool		`json:"Completed"`
}


func (table *Task) TableName() string {
	return "task"
}


func GetAllTasks(task *[]Task) (err error) {
	if err = Config.DB.Find(task).Error; err != nil {
		return err
	}

	return nil
}


func CreateTask(task *Task) (err error) {
	if err = Config.DB.Create(task).Error; err != nil {
		return err
	}

	return nil
}


func GetATask(task *Task, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(task).Error; err != nil {
		return err
	}

	return nil
}


func UpdateATask(task *Task) (err error) {
	Config.DB.Save(task)

	return nil
}


func DeleteATask(task *Task, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(task)

	return nil
}