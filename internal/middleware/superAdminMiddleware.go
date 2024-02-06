package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-newsApp-backend/pkg/jwtToken/adminToken"
	"log"
	"net/http"
)

func SuperAdminMiddleware(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	if token == "" {
		errResponse := response.Error(http.StatusUnauthorized, "Unauthorized", "Unauthorized", nil)
		return ctx.Status(fiber.StatusUnauthorized).JSON(errResponse)
	}
	claims, err := verifyAdminToken(token)
	if err != nil {
		errResponse := response.Error(http.StatusUnauthorized, "Invalid token", "Invalid token", nil)
		return ctx.Status(fiber.StatusUnauthorized).JSON(errResponse)
	}
	ctx.Locals("phone_number", claims.PhoneNumber)
	ctx.Locals("user_id", claims.AdminID)
	ctx.Locals("admin_role", claims.AdminRole)
	role := ctx.Locals("admin_role")
	log.Println(role)
	if role != "super_admin" {
		return fmt.Errorf("Permission denied")
	}
	return ctx.Next()
}

func verifyAdminToken(tokenString string) (*adminToken.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &adminToken.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return adminToken.SecretAdminKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*adminToken.Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("Invalid token")
}
