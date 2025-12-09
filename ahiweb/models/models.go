package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100);not null" binding:"required"`
	Password string `gorm:"type:varchar(25);not null" binding:"required"`
	Role     string `gorm:"type:varchar(20);not null" binding:"required"`
	Username string `gorm:"type:varchar(25);not null" binding:"required"`
}

type Posts struct {
	gorm.Model
	Header string `gorm:"type:varchar(100);not null" binding:"required"`
	Text   string `gorm:"type:varchar(500);not null" binding:"required"`
	Role   string `gorm:"type:varchar(20);not null" binding:"required"`
	Uid    uint
	Users  Users `gorm:"foreignKey:Uid"`
}

type Comments struct {
	gorm.Model
	Text  string `gorm:"type:varchar(500);not null" binding:"required"`
	Uid   uint
	Pid   uint
	Posts Posts `gorm:"foreignKey:Pid"`
	Users Users `gorm:"foreignKey:Uid"`
}
