package service

import (
	"time"

	"github.com/anjush-bhargavan/golib_admin/pkg/dom"
	"github.com/golang-jwt/jwt"
)

// GenerateToken to generate jwt token
func (a *AdminService) GenerateToken(userName string) (string, error) {
	admin, err := a.AdminRepo.FindAdmin(userName)
	if err != nil {
		return "", err
	}

	claims := &dom.Claims{
		AdminID: admin.AdminID,
		Email:  admin.Email,
		Role:   admin.Role,
		Claims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 10).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte("101101")

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
