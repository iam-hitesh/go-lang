package Services

import (
	"../Config"
	"../Models"
)

func GetAllTasks(task *[]Models.Task, ID uint) (err error) {
	if err = Config.DB.Where("created_by = ?", ID).Find(task).Error; err != nil {
		return err
	}

	return nil
}
