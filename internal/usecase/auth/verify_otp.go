package auth

import (
	"CleanArchitecture/internal/domain/model"
	"CleanArchitecture/internal/utils"
	"errors"
)

// فانکشن Verify OTP و ایجاد کاربر
func (u *AuthUsecase) VerifyOTPAndCreateUser(phone string, inputOtp string) (*model.User, error) {
	savedOtp, err := u.otpRepo.Get(phone)
	if err != nil {
		return nil, errors.New("otp not found or expired")
	}

	if savedOtp != inputOtp {
		return nil, errors.New("invalid otp")
	}

	_ = u.otpRepo.Save(phone, "", 1)

	// بررسی اینکه کاربر از قبل وجود دارد یا نه
	user, err := u.userRepo.FindByPhone(phone)
	if err != nil {
		// کاربر جدید
		user = &model.User{
			Phone: phone,
			Role:  utils.MakeRole(phone), // نقش همیشه ست می‌شود
		}

		if err := u.userRepo.Create(user); err != nil {
			return nil, err
		}
	}

	return user, nil
}
