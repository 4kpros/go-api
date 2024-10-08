package request

type CreateWithEmailRequest struct {
	Email string
	Role  int
}

type CreateWithPhoneNumberRequest struct {
	PhoneNumber int
	Role        int
}
