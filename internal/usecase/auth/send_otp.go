package auth

import (
	"CleanArchitecture/internal/utils/otp"
)

func (u *AuthUsecase) SendOTP(phone string) (string, error) {
	otp, err := otp.GenerateOTP()
	if err != nil {
		return "", err
	}

	// ذخیره در Redis (مثلاً 2 دقیقه)
	err = u.otpRepo.Save(phone, otp, 120)
	if err != nil {
		return "", err
	}

	return otp, nil
}
