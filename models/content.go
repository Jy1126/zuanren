package models

import "gorm.io/gorm"

type Content struct {
	gorm.Model
	ID          int64  `gorm:"id"`
	CN          string `gorm:"cn"`
	US          string `gorm:"us"`
	JP          string `gorm:"jp"`
	KR          string `gorm:"kr"`
	PH          string `gorm:"ph"`
	Other       string `gorm:"other"`
	OtherNote   string `gorm:"other_note"`
	AuditStatus int64  `gorm:"audit_status"`
}
