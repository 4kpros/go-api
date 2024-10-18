package mail

// Sends an email with the specified subject, message, and receiver.
func SendMail(subject string, message string, receiver string) error {
	return sendMailWithGmail(subject, message, receiver)
}
