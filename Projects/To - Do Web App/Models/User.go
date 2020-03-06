package Models

import (
	"github.com/jinzhu/gorm"
	//"golang.org/x/crypto/bcrypt"

	"../Config"
)


type User struct {
	gorm.Model

	Name string
	Email string 		`gorm:"type:varchar(100);unique_index;NOT NULL"`
	IsActive bool		`json:"is_active"`
	About string 		`json:"about";gorm:"type:text;"`
	Password string 	`json:"password";gorm:"NOT NULL"`
	Role int
	IsAdmin bool
}


func (user *User) TableName() string{
	return "user"
}

//func GetHashedPassword(password string) (string, error) {
//	bytes, err := bcrypt.GenerateFromPassword([]bytes(password), 14)
//
//	return string(bytes), err
//}
//
//
//func VerifyPassword(hashedPassword, rawPassword string) bool {
//	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
//
//	return err == nil
//}

func CreateUser(user *User) (err error) {
	// This will create a atomic transaction
	txn := Config.DB.Begin()

	err = txn.Create(&user).Error

	if err != nil {
		txn.Rollback()

		return err
	}

	txn.Commit()

	return nil
}


func GetAllUsers(users *[]User) (err error) {
	if err := Config.DB.Find(&users).Error; err != nil {
		return err
	}

	return nil
}


func GetAUser(user *User, id uint) (err error) {
	if err = Config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}

	return nil
}