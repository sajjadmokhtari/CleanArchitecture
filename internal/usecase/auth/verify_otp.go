package auth

import "errors"

func (u *AuthUsecase) VerifyOTP(phone string, inputOtp string) error {
	savedOtp, err := u.otpRepo.Get(phone)
	if err != nil {
		return errors.New("otp not found or expired")
	}

	if savedOtp != inputOtp {
		return errors.New("invalid otp")
	}

	// پاک کردن OTP بعد از موفقیت (مهم)
	_ = u.otpRepo.Save(phone, "", 1)

	return nil
}
