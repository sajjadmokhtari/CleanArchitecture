package jwt

import (
	"CleanArchitecture/internal/domain/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 🔹 این تابع یه توکن جدید می‌سازه برای کاربری که شماره تلفن و نقشش مشخصه
func GenerateJWT(userID uint, phone, role string) (string, error) {
	claims := model.CustomClaims{
		UserID: userID,
		Phone:  phone,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        GenerateJTI(),
			Issuer:    "your-app",
			Audience:  []string{"your-client"},
			Subject:   phone,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(privateKey)
}

// اعتبارسنجی JWT با کلید عمومی
func ValidateJWT(tokenStr string) (*model.CustomClaims, error) {
	claims := &model.CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
