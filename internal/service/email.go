package service

import (
	"net/smtp"

	"github.com/khairulharu/ewallet/domain"
	"github.com/khairulharu/ewallet/internal/config"
)

type emailService struct {
	cnf *config.Config
}

func NewEmail(cnf *config.Config) domain.EmailService {
	return &emailService{
		cnf: cnf,
	}
}

func (e emailService) Send(to, subject, body string) error {
	auth := smtp.PlainAuth("", e.cnf.Mail.User, e.cnf.Mail.Pass, e.cnf.Mail.Host)

	msg := []byte("From: khairulaswad <" + e.cnf.Mail.User + ">\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n" +
		body)

	return smtp.SendMail(e.cnf.Mail.Host+":"+e.cnf.Mail.Port, auth, e.cnf.Mail.User, []string{to}, msg)

}
