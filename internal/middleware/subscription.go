package middleware

import (
	config "NGB-SE/internal/conf"
	"NGB-SE/internal/util"
	"fmt"
	"time"
)

var emailTargets chan string

//
//发送用户的订阅邮件,每次发送后暂停一秒
//
func SendSubscription(targets <-chan string) {
	date := fmt.Sprintf("%d年%d月%d日", time.Now().Year(), time.Now().Month(), time.Now().Day())
	content := fmt.Sprintf("今天是%s，NGB-SE祝你有个好心情", date)
	for target := range targets {
		err := util.SendEmail(target, config.EmailConfig.EmailAddress, "text/plain", content)
		if err != nil {
			util.MakeInfoLog(fmt.Sprintf("Failed sending subscription email to %s,error:%s", target, err.Error()))
		} else {
			util.MakeInfoLog(fmt.Sprintf("Succeed in sending subscription email to %s", target))
		}
		time.Sleep(time.Second * time.Duration(1))
	}

}

func GetSubscriptor() {

}

func init() {
	emailTargets = make(chan string)
	for i := 0; i < 10; i++ {
		go SendSubscription(emailTargets)
	}
}
