package repository

import "CleanArchitecture/internal/domain/model"

type OTPRepository interface {// سیو و گت  او تی پی 
	Save(phone string, otp string, ttlSeconds int) error
	Get(phone string) (string, error)
}

type UserRepository interface { // ذخیره کاربر
	FindByPhone(phone string) (*model.User, error)
	Create(user *model.User) error
}
