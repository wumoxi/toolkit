package toolkit

import "github.com/go-gomail/gomail"

// 发送邮件配置参数
type MailerParams struct {
	ServerHost   string   // ServerHost 邮箱服务器地址，如腾讯邮箱为 smtp.qq.com
	ServerPort   int      // ServerPort 邮箱服务器端口，如腾讯邮箱为 465
	FromName     string   // 发件人名称
	FromEmail    string   // 发件人邮箱地址
	FromPassword string   // 发件人邮箱密码
	Toers        []string // 接收者邮件地址，可以添加多个
	CCers        []string // 抄送者邮件地址，可以添加多个
}

// 邮件结构体
type Mailer struct {
	message *gomail.Message
	params  *MailerParams
}

// 发送邮件
func (m *Mailer) Send(subject, body, contentType string) (bool, error) {
	m.message.SetHeader("Subject", subject)
	m.message.SetBody(contentType, body)
	dialer := gomail.NewDialer(m.params.ServerHost, m.params.ServerPort, m.params.FromEmail, m.params.FromPassword)
	if err := dialer.DialAndSend(m.message); err != nil {
		return false, err
	}
	return true, nil
}

// 获取一个邮件指针实例
func NewEmail(params *MailerParams) *Mailer {
	mailer := new(Mailer)
	mailer.message = gomail.NewMessage()
	mailer.params = params
	mailer.message.SetHeader("To", params.Toers...)
	mailer.message.SetHeader("Cc", params.CCers...)
	mailer.message.SetAddressHeader("From", params.FromEmail, params.FromName)
	return mailer
}
