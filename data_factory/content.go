package data_factory

import (
	"fmt"
	"zuanren/dao"
	"zuanren/models"
)

func GetContent(originCode string, targetCode string) []models.Content {
	var content []models.Content
	if originCode == "" {
		originCode = "cn"
	}
	sql := fmt.Sprintf("%s <> '' and audit_status = 101", originCode)
	condition := dao.DB.Where(sql)

	if targetCode != "" {
		sql = fmt.Sprintf("%s <> ''", targetCode)
		condition.Where(sql)
	}
	condition.Find(&content)

	return content
}

func AddContent(cn string, us string, jp string, kr string, ph string, other string, otherNote string) error {
	var content models.Content
	content.CN = cn
	content.US = us
	content.JP = jp
	content.KR = kr
	content.PH = ph
	content.Other = other
	content.OtherNote = otherNote
	content.AuditStatus = 100

	dao.DB.Create(&content)
	return nil
}
