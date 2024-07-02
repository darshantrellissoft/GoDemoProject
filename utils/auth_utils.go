package utils

import (
	"os"
	"strings"
	"time"

	"github.com/o1egl/paseto"
)

func GetPasetoSecretKey() []byte {
	return []byte(os.Getenv("PASETO_SECRET_KEY"))
}

var (
	pasetoV2 = paseto.NewV2()
)

func GenerateToken(userID string) (string, error) {
	now := time.Now()
	exp := now.Add(24 * time.Hour)
	nbt := now
	jsonToken := paseto.JSONToken{
		Audience:   "transaction-app",
		Issuer:     "transaction-app",
		Jti:        "123",
		Subject:    userID,
		IssuedAt:   now,
		Expiration: exp,
		NotBefore:  nbt,
	}

	footer := "some-footer"

	token, err := pasetoV2.Encrypt(GetPasetoSecretKey(), jsonToken, footer)
	if err != nil {
		return "", err
	}

	return token, nil
}

// ExtractToken extracts the token from the Authorization header
func ExtractToken(authHeader string) string {
	if authHeader == "" {
		return ""
	}

	// Check if the header has the format "Bearer <token>"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}

	return parts[1]
}
