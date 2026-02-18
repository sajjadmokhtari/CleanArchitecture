package auth

import "CleanArchitecture/internal/repository"

// لایه‌ی یوزکیس (منطق اصلی)
type AuthUsecase struct {
    otpRepo  repository.OTPRepository   // ریپوی OTP
    userRepo repository.UserRepository  // ریپوی کاربر
}

// سازنده‌ی یوزکیس
func NewAuthUsecase(otpRepo repository.OTPRepository, userRepo repository.UserRepository) *AuthUsecase {
    return &AuthUsecase{
        otpRepo:  otpRepo,   // تزریق OTP
        userRepo: userRepo,  // تزریق کاربر
    }
}

