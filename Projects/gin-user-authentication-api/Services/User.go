package Services


import (
	"../Config"
	"../Models"
)


func GetAllUsers(users *[]Models.User) (err error) {
	if err := Config.DB.Find(&users).Error; err != nil {
		return err
	}

	return nil
}