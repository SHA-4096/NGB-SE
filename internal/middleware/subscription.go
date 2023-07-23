package middleware

import (
	config "NGB-SE/internal/conf"
	"NGB-SE/internal/model"
	"NGB-SE/internal/util"
	"fmt"
	"time"
)

var emailTargets chan string

//
//发送用户的订阅邮件,每次发送后暂停一秒
//
func SendSubscription(targets <-chan string, id int) {
	util.MakeInfoLog(fmt.Sprintf("email goroutine %d has started", id))
	for {
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

}

//
//Send subscription email cronically
//
func CronSendScription() {
	for {
		if time.Now().Hour() == config.EmailConfig.SubscriptionHour {
			_, err := model.GetKeyValue("tag.subscriptionDone")
			if err != nil {
				users, err := model.QueryAllSubscriptors()
				if err != nil {
					util.MakeInfoLog("Error at middleware.CronSendScription")
					return
				} else {
					for _, user := range users {
						emailTargets <- user.Email
					}
				}
				//avoid sending email again
				model.SetKeyValuePair("tag.subscriptionDone", "subscriptionDone")
				model.SetExpiration("tag.subscriptionDone", 7200)
				util.MakeInfoLog("Finished subscription")
			}
		}
	}

}

func init() {
	emailTargets = make(chan string, 10)
	for id := 0; id < 10; id++ {
		go SendSubscription(emailTargets, id)
	}
	go CronSendScription()
}
