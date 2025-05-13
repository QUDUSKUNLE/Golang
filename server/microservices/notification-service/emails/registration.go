package emails

import (
	"context"
	"fmt"

	"github.com/QUDUSKUNLE/microservices/notification-service/templates"
	"gopkg.in/gomail.v2"
)

// SMTPConfig holds SMTP server configuration
type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

// SendRegistrationEmail sends a registration email and returns an error if it fails
func SendRegistrationEmail(ctx context.Context, smtpCfg SMTPConfig, name, email, token string) error {
	if name == "" || email == "" || token == "" {
		return fmt.Errorf("name, email, and token must not be empty")
	}
	emailBody := templates.RegistrationEmailTemplate(name, email, token)

	m := gomail.NewMessage()
	m.SetHeader("From", smtpCfg.From)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Verify Your Email Address")
	m.SetBody("text/html", emailBody)

	d := gomail.NewDialer(smtpCfg.Host, smtpCfg.Port, smtpCfg.Username, smtpCfg.Password)

	// Optionally, set a timeout using context
	done := make(chan error, 1)
	go func() {
	   done <- d.DialAndSend(m)
	}()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <- done:
		if err != nil {
			return fmt.Errorf("failed to send registration email: %w", err)
		}
	}
	return nil
}
