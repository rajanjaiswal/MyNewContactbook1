package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName   string `json:"fillname"`
	Email      string `json:"email"`
	UserName   string `json:"user_name"`
	Password   string `json:"password"`
	IsVerified bool   `json:"is_verified"`
}
