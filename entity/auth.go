package entity

import "time"

// RequestOTP request sent by the client
type RequestOTP struct {
	PhoneNumber string `json:"phone_number"`
}

// RequestOTPResponse is the reponse of otp request
type RequestOTPResponse struct {
	PhoneNumber string    `json:"phone_number"`
	SentTime    time.Time `json:"sent_time"`
}

// Login request sent from the client
type Login struct {
	PhoneNumber  string    `json:"phone_number"`
	OTPSentTime  time.Time `json:"otp_sent_time"`
	OTPCode      string    `json:"otp_code"`
	ClientUID    string    `json:"client_uid"`
	ClientSecret string    `json:"client_secret"`
}
