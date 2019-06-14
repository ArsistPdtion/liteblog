package models

import "github.com/jinzhu/gorm"

type Note struct {
	gorm.Model
	key string `gorm:"unique_index;not null"` //note unique identification
	UserID int //user id
	User User
	Title string //note title
	Summary string `gorm:"type:text"` //summary
	Content string `gorm:"type:text"`
	Visit int `gorm:"default:0"` //view count
	Praise int `gorm:"default:0"` //Thumb up number
}

func QueryNoteByKeyAndUserId(key string, userid int)(note Note, err error){
	return note, db.Model(&Note{}).Where("Key = ? and user_id = ?",key,userid).Take(&note).Error
}

func
