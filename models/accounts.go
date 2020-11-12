package models

import "gorm.io/gorm"

type XtpAccounts struct {
	gorm.Model
	ID         int64  `gorm:"id"`
	BindUserID int64  `gorm:"bind_user_id"`
	UserName   string `gorm:"username"`
	Password   string `gorm:"password"`
	Key        string `gorm:"key"`
}
