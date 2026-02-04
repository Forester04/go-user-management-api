package services

import (
	"fmt"
	"time"

	"github.com/Forester04/go-user-management-api/internal/errcode"
	"github.com/Forester04/go-user-management-api/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func (svc *Service) GenerateToken(user *models.User) (toeknString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Minute * 24 * 30).Unix(),
		"iat":   time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(viper.GetString("JWT_SECRET")))
	if err != nil {
		return "", fmt.Errorf("%w: %v", errcode.ErrGeneratingToken, err)
	}
	return tokenString, nil
}
