package handler

import (
	"CleanArchitecture/internal/usecase/auth"
	"CleanArchitecture/internal/utils/jwt"
	"CleanArchitecture/pkg/dto"
	"CleanArchitecture/pkg/validator"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// لایه‌ی هندلر (ورودی HTTP)
type AuthHandler struct {
	authUsecase *auth.AuthUsecase // یوزکیس احراز هویت
}

// سازنده‌ی هندلر
func NewAuthHandler(authUsecase *auth.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUsecase, // تزریق یوزکیس
	}
}

// SendOtpHandler godoc
// @Summary Send OTP
// @Description Send OTP to phone number
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.SendOtpRequest true "Phone number"
// @Success 200 {object} map[string]interface{}
// @Router /auth/send-otp [post]
func (h *AuthHandler) SendOtpHandler(c *gin.Context) {

	var req dto.SendOtpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if !validator.IsValidIranianMobile(req.Phone) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid mobile number"})
		return
	}

	resp := h.authUsecase.SendOTP(req.Phone)

	if !resp.Success {
		if resp.ErrType == "blocked" || resp.ErrType == "too_many" || resp.ErrType == "wait" {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": resp.Message})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": resp.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": resp.Message,
		"otp":     resp.Code,
	})

}

// VerifyOtpHandler godoc
// @Summary Verify OTP
// @Description Verify OTP code
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.VerifyOtpRequest true "Verify OTP"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/verify-otp [post]
func (h *AuthHandler) VerifyOtpHandler(c *gin.Context) {
	var req dto.VerifyOtpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := h.authUsecase.VerifyOTPAndCreateUser(req.Phone, req.OTP)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}

	role := user.Role

	token, err := jwt.GenerateJWT(user.ID, user.Phone, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"error":  "failed to generate token",
		})
		return
	}
	fmt.Println("token is :", token) //  برای تست

	c.SetCookie("access_token", token, 3600*24, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"user":   user,
		"role":   role,
	})
}
