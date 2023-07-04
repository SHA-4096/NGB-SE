package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Uid      string
	Name     string
	Password string
	JwtKey   string
	Email    string
	IsAdmin  string
}
