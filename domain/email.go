package domain

type EmailService interface {
	Send(to, subject, body string) error
}
