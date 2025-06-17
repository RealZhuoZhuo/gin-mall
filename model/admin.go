package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	UserName string `gorm:"type:varchar(50) not null"`
	Password string `gorm:"type:varchar(50) not null"`
	Avatar   string `gorm:"size:255"`
}
