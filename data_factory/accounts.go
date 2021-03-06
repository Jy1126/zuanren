package data_factory

import (
	"zuanren/dao"
	"zuanren/models"
)

func GetBindCount() int64 {
	var count int64
	dao.DB.Where("bind_user_id > ?", 0).Find(&models.XtpAccounts{}).Count(&count)
	return count
}

func GetUnBindCount() int64 {
	var count int64
	dao.DB.Where("bind_user_id = ?", 0).Find(&models.XtpAccounts{}).Count(&count)
	return count
}
