package auth

import (
	"CleanArchitecture/internal/domain/model"
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

	// پاک کردن OTP بعد از موفقیت
	_ = u.otpRepo.Save(phone, "", 1)

	//  ساخت توکن  و سیو توی کوکی ها 

	// بررسی اینکه کاربر از قبل وجود دارد یا نه
	user, err := u.userRepo.FindByPhone(phone)
	if err != nil {
		user = &model.User{Phone: phone}
		if err := u.userRepo.Create(user); err != nil {
			return nil, err
		}
	}

	return user, nil
}
