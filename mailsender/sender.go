package mailsender

import (
	"net/smtp"
)

// Mail struct have attribute to need to sending mail
type Mail struct {
	From     string
	Username string
	Password string
	To       string
	Sub      string
	Msg      string
}

// Body creates body of mail by Mail's attribute
func (m Mail) Body() string {
	return "To: " + m.To + "\r\n" +
		"Subject: " + m.Sub + "\r\n\r\n" +
		m.Msg + "\r\n"
}

// SendGmail send to user who indicated Mail'sattribute by gmail
func SendGmail(m Mail) error {
	smtpSvr := "smtp.gmail.com:587"
	auth := smtp.PlainAuth("", m.Username, m.Password, "smtp.gmail.com")
	if err := smtp.SendMail(smtpSvr, auth, m.From, []string{m.To}, []byte(m.Body())); err != nil {
		return err
	}
	return nil
}