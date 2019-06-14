package models

import (
	"github.com/jinzhu/gorm"
	"os"
)

var db gorm.DB


func init(){
	//create data directory
	if err := os.MkdirAll("data",0777);err!=nil{
		panic("failed to create database directory,"+err.Error())
	}
	//open database, copy to db variabe, file patyh is data/data.db
	db, err := gorm.Open("sqlite3","data/data.db")
	if err != nil{
		panic("failed to connect database,"+err.Error())
	}
	//automatic synchronization of table structures
	db.AutoMigrate(&User{})
	var count int
	//Model(&User{}) search user table, Count(&count) assign data to the count field.
	if err := db.Model(&User{}).Count(&count).Error; err == nil && count ==0{
		//add admin user
		db.Create(&User{
			Name:"admin",
			Email:"admin@qq.com",
			Pwd:"123123",
			Avatar:"/static/images/info-img.png",
			Role:0,
		})
	}
}

