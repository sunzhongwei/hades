package hades

import (
	//gomail "gopkg.in/gomail.v2"  // 改用 https://github.com/Shopify/gomail
	"regexp"

	gomail "github.com/Shopify/gomail"
)

type SMTPConfig struct {
	Host     string
	Port     int
	User     string
	Pass     string
	UserName string // 发件人名称
}

// SendMail 发送邮件
// 配置参数示例：
/*
{
	"Host": "smtp.163.com",
	"Port": 465,
	"User": "user@163.com",
	"Pass": "password",
	"UserName": "发件人名称"
}
*/
func SendMail(to, cc, subject, body string, config SMTPConfig) error {
	// 解析配置
	/*
		smtpConfig := cache.CacheGet("smtp_config")
		var config SMTPConfig
		smtpConfigBytes, ok := smtpConfig.([]byte)
		if !ok {
			return fmt.Errorf("smtp_config is not of type []byte")
		}
		if err := json.Unmarshal(smtpConfigBytes, &config); err != nil {
			return err
		}
	*/

	// 创建邮件消息
	m := gomail.NewMessage()
	m.SetHeader("From", config.User)
	m.SetHeader("To", to)
	if cc != "" {
		m.SetHeader("Cc", cc)
	}
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	// 设置邮件服务器
	d := gomail.NewDialer(config.Host, config.Port, config.User, config.Pass)

	// DialAndSend 的超时是通过 Dialer.Timeout 字段进行设置的
	// d.Timeout = 30 * time.Second // 设置为 30 秒。默认是 10 秒

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

// email address validation regex
func IsValidEmail(email string) bool {
	// simplified regex for demonstration purposes
	const emailRegex = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
