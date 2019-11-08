package toolkit

import (
	"testing"
)

func TestMailer_Send(t *testing.T) {
	mailer := NewEmail(&MailerParams{
		ServerHost:   "smtp.qq.com",
		ServerPort:   465,
		FromEmail:    "wu.shaohua@foxmail.com",
		FromPassword: "mmooqssdsuwjskssddthpubddf",
		FromName:     "武沫汐",
		Toers:        []string{"warnerwu@126.com", "warnerwu@139.com", "contact.shaohua@gmail.com"},
		CCers:        []string{"warnerwu@163.com"},
	})

	send, err := mailer.Send("Golang邮件发送", `中华人民共和国 - Golang邮件发送`, "text/plain")
	if err != nil && !send {
		t.Fatal(err)
	} else {
		t.Log("发送成功")
	}
}
