package Models

import (
	"github.com/jinzhu/gorm"

	"../Config"
	"../Helpers"
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


type UserInterface interface {
	TableName() string
	Create() (err error)
	GetByEmail(email string) (err error)
}

func (user *User) TableName() string{
	return "user"
}


func (user *User) Create() (err error) {
	// This will create a atomic transaction
	txn := Config.DB.Begin()

	// Update raw password to Hashed password before storing to DB
	user.Password, _ = Helpers.GetHashedPassword(user.Password)

	err = txn.Create(&user).Error

	if err != nil {
		txn.Rollback()

		return err
	}

	txn.Commit()

	return nil
}


func (user *User) GetByEmail(email string) (err error) {
	if err = Config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return err
	}

	return nil
}