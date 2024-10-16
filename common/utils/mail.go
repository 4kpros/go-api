package utils

import (
	"fmt"
	"net/smtp"

	"api/config"
)

// Sends an email with the specified subject, message, and receiver.
func SendMail(subject string, message string, receiver string) error {
	return sendMailWithGmail(subject, message, receiver)
}

// Sends an email using Gmail's SMTP server
func sendMailWithGmail(subject string, message string, receiver string) error {
	to := []string{receiver}
	msg := []byte(
		fmt.Sprintf(
			"To: %s\r\n"+
				"Subject: Go-api %s\r\n"+
				"\r\n"+
				"%s\r\n",
			receiver, subject, message,
		),
	)

	return smtp.SendMail(
		fmt.Sprintf("%s:%d", config.Env.SmtpHost, config.Env.SmtpPort),
		config.LoginSMTPAuth(config.Env.SmtpUsername, config.Env.SmtpPassword),
		config.Env.SmtpSender,
		to,
		msg,
	)
}
