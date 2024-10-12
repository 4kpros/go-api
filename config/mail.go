package config

import (
	"errors"
	"fmt"
	"net/smtp"
)

type SMTPAuth struct {
	Username string
	Password string
}

// Creates an SMTP authentication mechanism using
// the provided username and password.
func LoginSMTPAuth(username, password string) smtp.Auth {
	return &SMTPAuth{
		Username: username,
		Password: password,
	}
}

// Initiates the SMTP authentication process by sending the "LOGIN" command.
// Returns the "LOGIN" command, an empty byte slice, and no error.
func (a *SMTPAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

// Processes the server's response during SMTP authentication.
// If 'more' is true, sends the username or password based on the server's challenge.
// Returns the next client response and an error if any occurs.
func (a *SMTPAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.Username), nil
		case "Password:":
			return []byte(a.Password), nil
		default:
			return nil, errors.New("unknown from server")
		}
	}
	return nil, nil
}

// Sends an email with the specified subject, message, and receiver.
func SendMail(subject string, message string, receiver string) error {
	return SendWithGmail(subject, message, receiver)
}

// Sends an email using Gmail's SMTP server
func SendWithGmail(subject string, message string, receiver string) error {
	var to = []string{receiver}
	var msg = []byte(
		fmt.Sprintf(
			"To: %s\r\n"+
				"Subject: Go-api %s\r\n"+
				"\r\n"+
				"%s\r\n",
			receiver, subject, message,
		),
	)

	return smtp.SendMail(
		fmt.Sprintf("%s:%d", Env.SmtpHost, Env.SmtpPort),
		LoginSMTPAuth(Env.SmtpUsername, Env.SmtpPassword),
		Env.SmtpSender,
		to,
		msg,
	)
}
