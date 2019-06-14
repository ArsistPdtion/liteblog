package models

import (
	"github.com/jinzhu/gorm"
)



type User struct {
	gorm.Model
	Name string `gorm:"unique_index"`
	Email string `gorm:"unique_index"`
	Avatar string
	Pwd string
	Role int `gorm:"default:1"` //0 administrator 1 common user
}

func QueryUserByEmailAndPassword(email,password string)(user User, err error){
	return user, db.Model(&User{}).Where("email = ? and pwd = ?",email,password).Take(&user).Error
}

//query user by name
func QueryUserByName(name string)(user User, err error){
	return user, db.Where("name = ?", name).Take(&user).Error
}

//query user by email
func QueryUserByEmail(email string)(user User, err error){
	return user, db.Where("email = ?",email).Take(&user).Error
}

//save user
func SaveUser(user *User)error{
	return db.Create(user).Error
}


