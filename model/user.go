package model

import "github.com/zhime/monitor/global"

type User struct {
	global.MODEL
	Name     string `gorm:"size:100"`
	Email    string `gorm:"size:100;unique"`
	Password string `gorm:"size:255"`
}

func (User) TableName() string {
	return "user"
}
