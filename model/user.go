package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null;unique"`
	Telephone string `gorm:"type:varchar(110);not null"`
	Password  string `gorm:"type:varchar(255);not null"`
}
