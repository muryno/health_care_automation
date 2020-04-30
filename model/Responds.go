package model



type RegistrationResponds struct {
	ID           uint    `json:"id"`
	Email        string `json:"email"`
	Token        string `json:"Authorization"`
	OTP        string `json:"otp_code"`
	Status        uint   `json:"user_status"`

}