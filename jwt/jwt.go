package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/arezooq/open-utils/errors"
)

type CustomClaims struct {
    UserID uuid.UUID
}

func GenerateAccessToken(userID uuid.UUID) (string, error) {
	err := godotenv.Load(".env")

	if err != nil {
		return "", errors.ErrInternal
	}

	JWTSecret := os.Getenv("SECRET_JWT")

	claims := jwt.RegisteredClaims{
		Subject:   userID.String(),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour).UTC()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(JWTSecret))
	if err != nil {
		return "", errors.ErrInternal
	}

	return tokenString, nil
}

func GenerateRefreshToken(userID uuid.UUID) (string, error) {
	err := godotenv.Load(".env")

	if err != nil {
		return "", errors.ErrInternal
	}

	JWTSecret := os.Getenv("SECRET_JWT")

	claims := jwt.RegisteredClaims{
		Subject:   userID.String(),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour).UTC()), // 7 days
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(JWTSecret))
	if err != nil {
		return "", errors.ErrInternal
	}

	return tokenString, nil
}

func ValidateRefreshToken(tokenStr string) (*CustomClaims, error) {
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("REFRESH_SECRET")), nil
    })
    if err != nil || !token.Valid {
        return nil, err
    }

    claims := token.Claims.(jwt.MapClaims)
    return &CustomClaims{
        UserID: uuid.MustParse(claims["user_id"].(string)),
    }, nil
}