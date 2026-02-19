package jwt

import (
	"crypto/rsa"
)

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

// بارگذاری کلیدها هنگام راه‌اندازی سرور
func InitJWTKeys(privatePath, publicPath string) error {
	var err error
	privateKey, err = LoadPrivateKey(privatePath)
	if err != nil {
		return err
	}
	publicKey, err = LoadPublicKey(publicPath)
	if err != nil {
		return err
	}
	return nil
}
