package service

import (
	df "zuanren/data_factory"
	"zuanren/models"
)

func GetContent(originCode string, targetCode string) []models.Content {
	return df.GetContent(originCode, targetCode)
}

func AddContent(cn string, us string, jp string, kr string, ph string, other string, otherNote string) error {
	return df.AddContent(cn, us, jp, kr, ph, other, otherNote)
}
