package modelerr

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       int    `json:"id" gorm:"primary_key;autoIncrement:true"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
