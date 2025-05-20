package infra

import (
	"fmt"
	"msvc-users/application/contracts"
	"net/smtp"
)

type GmailMailer struct {
	Username string
	Password string
}

func NewGmailMailer(username, password string) contracts.EmailSender {
	return &GmailMailer{Username: username, Password: password}
}

var _ contracts.EmailSender = (*GmailMailer)(nil)

func (m *GmailMailer) Send(to, subject, body string) error {
	from := m.Username
	pass := m.Password

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	headers := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n", from, to, subject)

	auth := smtp.PlainAuth("", from, pass, smtpHost)

	message := []byte(headers + body)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
