package mailcore

import (
	"bytes"
	"html/template"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

//Params is used as a params and value
type Params map[string]interface{}

// New is used to set a new mail
type New struct {
	To []string
	Cc []struct {
		Addr string
		Name string
	}
	Bcc        []string
	Subject    string
	Params     Params
	Attachment string
}

// MailConf is
const MailConf = "./template/mailconf.html"

// Send is used to send after object has been set
func (m *New) Send(tmpl string) error {
	var err error
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		return err
	}
	t, err = t.ParseFiles(tmpl)
	if err != nil {
		return err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, m.Params); err != nil {
		return err
	}

	result := tpl.String()
	mail := gomail.NewMessage()
	mail.SetHeader("From", os.Getenv("MAIL_FROM"))
	mail.SetHeader("To", m.To...)
	// mail.SetAddressHeader("Cc", "<RECIPIENT CC>", "<RECIPIENT CC NAME>")
	mail.SetHeader("Subject", m.Subject)
	mail.SetBody("text/html", result)
	if m.Attachment != "" {
		mail.Attach(m.Attachment)
	}
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	d := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		port,
		os.Getenv("SMTP_EMAIL"),
		os.Getenv("SMTP_PASSWORD"),
	)

	if err := d.DialAndSend(mail); err != nil {
		return err
	}

	return nil
}
