package config

import (
	"errors"
	"fmt"
	"net/smtp"
)

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("unknown from server")
		}
	}
	return nil, nil
}

func SendMail(subject string, message string, receiver string) (err error) {
	err = SendWithGmail(subject, message, receiver)
	return
}

func SendWithGmail(subject string, message string, receiver string) (err error) {
	// Choose auth method and set it up
	auth := LoginAuth(AppEnv.SmtpUsername, AppEnv.SmtpPassword)

	// Here we do it all: connect to our server, set up a message and send it
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
	err = smtp.SendMail(
		fmt.Sprintf("%s:%d", AppEnv.SmtpHost, AppEnv.SmtpPort),
		auth,
		AppEnv.SmtpSender,
		to,
		msg,
	)

	return
}
