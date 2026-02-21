package repository

type RateLimiterRepository interface {
    CanSendOTP(phone string) (bool, error)
    MarkOTPSent(phone string) error
    OTPRequestCount(phone string) (int, error)
    IncrementOTPRequest(phone string) error
    GetFailedAttempts(phone string) (int, error)
    IncrementFailedAttempts(phone string) error
}
