package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primarykey"`
	Username string `json:"username" form:"username" query:"username"`
	Password string `json:"password" form:"password" query:"password"`
}

func CreateUser(db *gorm.DB, User *User) (err error) {
	err = db.Table("users").Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserById(db *gorm.DB, User *User, id int) (err error) {
	err = db.Table("users").Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(db *gorm.DB, User *User) (err error) {
	db.Table("users").Save(User)
	return nil
}

func DeleteUser(db *gorm.DB, User *User, id int) (err error) {
	db.Table("users").Where("id = ?", id).Delete(User)
	return nil
}
