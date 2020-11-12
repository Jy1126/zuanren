package service

import (
	"fmt"
	"strconv"
	df "xtp_account_check/data_factory"
	"xtp_account_check/utils"
)

func GetBindCount() (int64, int64) {
	bindCount := df.GetBindCount()
	unBindCount := df.GetUnBindCount()
	return bindCount, unBindCount
}

func SendCheckCount() {
	unBindCount := df.GetUnBindCount()
	if unBindCount < 50 {
		title := fmt.Sprintf("xtp测试账户目前剩余 %s 个  请处理", strconv.FormatInt(unBindCount, 10))
		_ = utils.SendMail(utils.Config.SendAddress, title, "")
	}

}
