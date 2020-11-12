package service

import (
	df "xtpwidget/data_factory"
)

func GetBindCount() (int64, int64) {
	bindCount := df.GetBindCount()
	unBindCount := df.GetUnBindCount()
	return bindCount, unBindCount
}
