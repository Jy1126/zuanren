package utils

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func SendMail(to string, title, content string) error {

	requestUrl := Config.SendmailUrl

	postValue := url.Values{
		"cc":       {""},
		"from":     {"xtp@zts.com.cn"},
		"mailtype": {"html"},
		"to":       {to},
		"sender":   {"xtp@zts.com.cn"},
		"title":    {title},
		"content":  {content},
	}
	request, err := http.PostForm(requestUrl, postValue)
	if err != nil {
		return err
	}
	defer request.Body.Close()

	if request.StatusCode == 200 {
		rb, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return err
		}
		log.Println("发送邮件成功:", string(rb))
		return nil
	}

	return errors.New("发送邮件失败")
}
