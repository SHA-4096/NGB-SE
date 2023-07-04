package util

import (
	config "NGB-SE/internal/conf"
	"fmt"

	"gopkg.in/gomail.v2"
)

var (
	emailClient *gomail.Dialer
)

//
//初始化邮件的Client，需要在读取完配置之后
//
func EmailClientInit() {
	emailClient = gomail.NewDialer(config.EmailConfig.SmtpServer, config.EmailConfig.SmtpPort, config.EmailConfig.Name, config.EmailConfig.Password)
	//emailClient.TLSConfig = &tls.Config{InsecureSkipVerify: true} //注意：用于在主机没有有效证书的情况下使用，不可用于生产环境
}

//
//发送邮件，ContentType一般为text/html
//
func SendEmail(TargetEmailAddress, Header, ContentType, Content string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", config.EmailConfig.EmailAddress)
	m.SetHeader("To", TargetEmailAddress)
	//	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", Header)
	m.SetBody("text/html", Content)
	MakeInfoLog(fmt.Sprintf("使用账户%s向%s发送邮件，标记发送者为%s，端口号为%d", config.EmailConfig.Name, TargetEmailAddress, config.EmailConfig.EmailAddress, config.EmailConfig.SmtpPort))
	if err := emailClient.DialAndSend(m); err != nil {
		MakeInfoLog(fmt.Sprintf("哦莫，邮件没有正常发送,错误信息：%s", err.Error()))
		return err
	}
	return nil

}
