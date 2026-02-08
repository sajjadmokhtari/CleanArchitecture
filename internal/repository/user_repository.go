package repository



 

type OTPRepository interface {
	Save(phone string, otp string, ttlSeconds int) error
	Get(phone string) (string, error)
}
