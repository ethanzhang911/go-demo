package main

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/gomail.v2"
)

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "21802259@qq.com")
	m.SetHeader("To", "zhangcheng@cj.com.cn")
	m.SetHeader("Subject", "这是一份通过gomail测试邮件")
	m.SetBody("text/plain", "邮件正文")
	d := gomail.NewDialer("smtp.qq.com", 465, "21802259@qq.com", "neqfwlxdbymobibe")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", "发送完成")
}
