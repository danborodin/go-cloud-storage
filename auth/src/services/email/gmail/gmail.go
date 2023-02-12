package gmail

import (
	"auth/src/configs"
	"auth/src/repository/database"
	"github.com/danborodin/go-logd"
	"net/smtp"
)

type Service struct {
	l  *logd.Logger
	db database.DB
}

//var smtpHostname = "smtp.gmail.com:587"

func NewGmailService(l *logd.Logger, db database.DB) *Service {
	return &Service{
		l:  l,
		db: db,
	}
}

func (s Service) SendEmail(subject, text string, to []string, contentType string) error {
	mime := "MIME-version: 1.0;\nContent-Type: " + contentType + "; charset=\"UTF-8\";\n\n"
	subj := "Subject: " + subject + "!\n"
	msg := []byte(subj + mime + "\n" + text)
	auth := smtp.PlainAuth("", configs.Conf.Gmail.Username, configs.Conf.Gmail.Pwd, configs.Conf.Gmail.Host)

	err := smtp.SendMail(configs.Conf.Gmail.Host+":"+configs.Conf.Gmail.Port, auth, configs.Conf.Gmail.Username, to, msg)
	if err != nil {
		return err
	}

	s.l.InfoPrintln("email sent")

	return nil
}
