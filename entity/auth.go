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
