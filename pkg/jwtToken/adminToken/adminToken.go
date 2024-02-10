package adminToken

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type AdminClaims struct {
	AdminID     int    `json:"admin_id"`
	PhoneNumber string `json:"phone_number"`
	AdminRole   string `json:"admin_role"`
	jwt.StandardClaims
}

var SecretAdminKey = []byte("das#jd!ahDjSwr$we$ry$wbw_we^t*&^$%^#$sa)adEd&$sda")

func GenerateAdminToken(adminID int, phoneNumber, adminRole string) (string, error) {
	adminClaims := AdminClaims{
		AdminID:     adminID,
		PhoneNumber: phoneNumber,
		AdminRole:   adminRole,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, adminClaims)
	tokenString, err := token.SignedString(SecretAdminKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
