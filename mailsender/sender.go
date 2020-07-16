package mailsender

import (
    "net/smtp"
)

// Mail struct have attribute to need to sending mail
type Mail struct{
	from string
	username string
	password string
	to string
	sub string
	msg string
}

// BuildMail : constructer of Mail
func BuildMail(from,username,password,to,sub,msg string) Mail{
	return Mail{from,username,password,to,sub,msg}
}

// Body creates body of mail by Mail's attribute
func (m Mail) Body() string {
    return "To: " + m.to + "\r\n" +
        "Subject: " + m.sub + "\r\n\r\n" +
        m.msg + "\r\n"
}

// SendGmail send to user who indicated Mail'sattribute by gmail
func SendGmail(m Mail) error {
    smtpSvr := "smtp.gmail.com:587"
    auth := smtp.PlainAuth("", m.username, m.password, "smtp.gmail.com")
    if err := smtp.SendMail(smtpSvr, auth, m.from, []string{m.to}, []byte(m.Body())); err != nil {
        return err
    }
    return nil
}