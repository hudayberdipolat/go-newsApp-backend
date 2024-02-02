package adminToken

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type Claims struct {
	AdminID     int    `json:"admin_id"`
	PhoneNumber string `json:"phone_number"`
	AdminRole   string `json:"admin_role"`
	jwt.StandardClaims
}

var SecretKey = []byte("das#jd!ahDjSwr$we$ry$wbw_we^t*&^$%^#$sa)adEd&$sda")

func GenerateAdminToken(adminID int, phoneNumber, adminRole string) (string, error) {
	claims := Claims{
		AdminID:     adminID,
		PhoneNumber: phoneNumber,
		AdminRole:   adminRole,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
