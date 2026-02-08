package auth

import (
	"CleanArchitecture/internal/repository"
	"CleanArchitecture/internal/utils/otp"
)

type AuthUsecase struct {
	otpRepo repository.OTPRepository
}

func NewAuthUsecase(otpRepo repository.OTPRepository) *AuthUsecase {
	return &AuthUsecase{
		otpRepo: otpRepo,
	}
}




func (u *AuthUsecase) SendOTP(phone string) (string, error) {
	code, err := otp.GenerateOTP()
	if err != nil {
		return "", err
	}

	// ذخیره در Redis (مثلاً 2 دقیقه)
	err = u.otpRepo.Save(phone, code, 120)
	if err != nil {
		return "", err
	}

	return code, nil
}
