package auth

import "CleanArchitecture/internal/infrastructure/auth"

type AuthUsecase struct {}

func NewAuthUsecase() *AuthUsecase {
	return &AuthUsecase{}
}

// SendOTP فقط OTP تولید می‌کند
func (u *AuthUsecase) SendOTP(phone string) string {
	otp := auth.GenerateOTP()
	// اینجا بعداً می‌تونیم Redis اضافه کنیم
	return otp
}
