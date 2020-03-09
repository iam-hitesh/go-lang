package Models

import (
	"github.com/jinzhu/gorm"

	"../Config"
)


type Task struct {
	gorm.Model

	Title string 		`json:"Title"`
	Description string 	`json:"Description"`
	Completed bool		`json:"Completed"`
	User User			`gorm:"foreignkey:Id; NOT NULL;"`
	CreatedBy uint
}


type TaskInterface interface {
	TableName() string
	Create() (err error)
	Get(id string) (err error)
	Update() (err error)
	Delete(id string) (err error)
}


func (task *Task) TableName() string {
	return "task"
}


func (task *Task) Create() (err error) {
	if err = Config.DB.Create(task).Error; err != nil {
		return err
	}

	return nil
}


func (task *Task) Get(id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(task).Error; err != nil {
		return err
	}

	return nil
}


func (task *Task) Update() (err error) {
	Config.DB.Save(task)

	return nil
}


func (task *Task) Delete(id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(task)

	return nil
}


func GetAllTasks(task *[]Task) (err error) {
	if err = Config.DB.Find(task).Error; err != nil {
		return err
	}

	return nil
}