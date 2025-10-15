package jwt

import (
	"os"
	"strings"
	"time"

	"github.com/arezooq/open-utils/api"
	"github.com/arezooq/open-utils/errors"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// Generate Access Token
func GenerateAccessToken(userID string) (string, error) {
	secret := os.Getenv("SECRET_JWT")
	claims := CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// Generate Refresh Token
func GenerateRefreshToken(userID string) (string, error) {
	secret := os.Getenv("REFRESH_SECRET")
	claims := CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// Validate Refresh Token
func ValidateRefreshToken(tokenStr string) (*CustomClaims, error) {
	claims := &CustomClaims{}
	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}

// Extract User ID From Token
func ExtractUserIDFromToken(tokenStr string) (string, error) {
	claims := &CustomClaims{}
	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_JWT")), nil
	})
	if err != nil {
		return "", err
	}
	return claims.UserID, nil
}

// ExtractTokenFromHeader
func ExtractTokenFromHeader(req *api.Request) (string, error) {
	authHeader := req.Ctx.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.ErrMissingToken
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.ErrInvalidToken
	}
	req.Token = parts[1]

	return req.Token, nil
}
