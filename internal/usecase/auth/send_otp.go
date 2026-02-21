package auth

import (
	"CleanArchitecture/internal/utils/otp"
)

func (u *AuthUsecase) SendOTP(phone string) OTPResponse {

	failed, _ := u.rateLimiter.GetFailedAttempts(phone)
	if failed >= 5 {
		return OTPResponse{
			Success: false,
			ErrType: "blocked",
			Message: "کاربر موقتاً بلاک شده است",
		}
	}

	canSend, _ := u.rateLimiter.CanSendOTP(phone)
	if !canSend {
		return OTPResponse{
			Success: false,
			ErrType: "wait",
			Message: "لطفاً چند ثانیه صبر کنید",
		}
	}

	count, _ := u.rateLimiter.OTPRequestCount(phone)
	if count >= 5 {
		return OTPResponse{
			Success: false,
			ErrType: "too_many",
			Message: "تعداد درخواست بیش از حد مجاز است",
		}
	}

	code, err := otp.GenerateOTP()
	if err != nil {
		return OTPResponse{
			Success: false,
			ErrType: "internal",
			Message: "خطای داخلی",
		}
	}

	err = u.otpRepo.Save(phone, code, 120)
	if err != nil {
		return OTPResponse{
			Success: false,
			ErrType: "internal",
			Message: "خطای ذخیره‌سازی OTP",
		}
	}

	u.rateLimiter.MarkOTPSent(phone)
	u.rateLimiter.IncrementOTPRequest(phone)

	return OTPResponse{
		Success: true,
		Code:    code,
		Message: "OTP ارسال شد",
	}
}
