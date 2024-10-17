package sms

const unavailableErrMsg = "is not available! Please try again later."
const emptyMessageErrMsg = "Empty message! Please add message content."
const emptyReceiverErrMsg = "Empty receiver! Please add receiver."
const emptyRecipientsErrMsg = "Empty recipients! Please add least one receiver."

// Sends a SMS to a specified receiver
func SendSMS(message string, receiver string) error {
	return sendSMSWithTwilio(message, receiver)
}

// Send a SMS to multiple recipients
func SendBulkSMS(message string, recipients []string) error {
	return sendBulkSMSWithTwilio(message, recipients)
}
