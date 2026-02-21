package auth

type OTPResponse struct {
    Success bool
    Code    string
    ErrType string
    Message string
}
