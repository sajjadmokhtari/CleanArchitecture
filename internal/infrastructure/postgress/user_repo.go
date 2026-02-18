package postgres

import (
	"CleanArchitecture/internal/domain/model"
	"CleanArchitecture/internal/repository"

	"gorm.io/gorm"
)

type UserPostgresRepository struct {
	db *gorm.DB
}

func NewUserPostgresRepository(db *gorm.DB) repository.UserRepository {
	return &UserPostgresRepository{db: db}
}





func (r *UserPostgresRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}





func (r *UserPostgresRepository) FindByPhone(phone string) (*model.User, error) {
	var user model.User
	err := r.db.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

