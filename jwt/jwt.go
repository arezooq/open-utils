package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
    UserID string `json:"user_id"`
    jwt.RegisteredClaims
}

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
