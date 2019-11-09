package toolkit

import (
	"testing"
)

func TestMailer_Send(t *testing.T) {
	mailer := NewEmail(&MailerParams{
		ServerHost:   "smtp.qq.com",
		ServerPort:   465,
		FromEmail:    "shaohua@foxmail.com",
		FromPassword: "mmooqssdsuwjskssddthpubddf",
		FromName:     "武沫汐",
		Toers:        []string{"warner@126.com", "warner@139.com", "contact.warner@gmail.com"},
		CCers:        []string{"warner@163.com"},
	})

	send, err := mailer.Send("Golang邮件发送", `中华人民共和国 - Golang邮件发送`, "text/plain")
	if err != nil && !send {
		t.Fatal(err)
	} else {
		t.Log("发送成功")
	}
}
