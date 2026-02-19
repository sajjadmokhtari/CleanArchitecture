package jwt

import (
	"github.com/google/uuid"
)

func GenerateJTI() string {
	return uuid.New().String()
}
