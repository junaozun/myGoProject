package model

import (
	"weblogindemo/encry"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20"`
}

// 钩子函数，用于加密密码
func (u *User) BeforeSave() {
	u.Password = encry.ScryptPasswd(u.Password)
}
