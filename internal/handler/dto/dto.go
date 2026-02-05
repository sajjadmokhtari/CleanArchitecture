package dto

type SendOtpRequest struct {
    Phone string `json:"phone"`
}

type VerifyOtpRequest struct {
    Phone string `json:"phone"`
    OTP   string `json:"otp"`
}
