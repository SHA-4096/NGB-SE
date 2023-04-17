package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Uid      string
	Name     string
	Password string
	Jwt_key  string
}
