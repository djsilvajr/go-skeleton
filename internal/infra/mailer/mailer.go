package mailer

import (
	"fmt"
	"net/smtp"

	"github.com/djsilvajr/go-skeleton/internal/config"
)

type Mailer struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Mailer {
	return &Mailer{cfg: cfg}
}

// Send sends a plain-text email. Extend to support HTML templates as needed.
func (m *Mailer) Send(to []string, subject, body string) error {
	auth := smtp.PlainAuth("", m.cfg.MailUser, m.cfg.MailPassword, m.cfg.MailHost)

	msg := fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/plain; charset=utf-8\r\n\r\n%s",
		m.cfg.MailFrom, to[0], subject, body,
	)

	addr := fmt.Sprintf("%s:%s", m.cfg.MailHost, m.cfg.MailPort)
	return smtp.SendMail(addr, auth, m.cfg.MailFrom, to, []byte(msg))
}
