/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/5 15:30
 * @File:  SendEmail
 * @Software: GoLand
 **/

package utils

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

// SendEmail 发送邮箱
func SendEmail(toUserEmail, code string) error {
	e := email.NewEmail()               // 创建邮件对象
	e.From = "Code <2865260231@qq.com>" // 发送人
	e.To = []string{toUserEmail}        // 接收人
	e.Subject = "验证码已经发送"               // 邮件主题
	e.HTML = []byte("您的验证码：<b>" + code + "</b>")
	err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", "2865260231@qq.com", "hznbdiagkbqedeaa", "smtp.qq.com")) // 发送邮件
	if err != nil {
		fmt.Println("发送邮件失败", err)
		return err
	}
	return nil
}
