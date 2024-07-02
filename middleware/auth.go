package middleware

import (
	"MyTransactAPP/utils"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/o1egl/paseto"
)

// var secretKey = []byte(os.Getenv("PASETO_SECRET_KEY"))
func GetPasetoSecretKey() []byte {
	return []byte(os.Getenv("PASETO_SECRET_KEY"))
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.JSONResponse(c, http.StatusUnauthorized, "Unauthorized", gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := utils.ExtractToken(authHeader)
		if tokenString == "" {
			utils.JSONResponse(c, http.StatusUnauthorized, "Unauthorized", gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		var jsonToken paseto.JSONToken
		var footer string

		err := paseto.NewV2().Decrypt(tokenString, utils.GetPasetoSecretKey(), &jsonToken, &footer)
		if err != nil {
			utils.JSONResponse(c, http.StatusUnauthorized, "Unauthorized", gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if jsonToken.Expiration.Before(time.Now()) {
			utils.JSONResponse(c, http.StatusUnauthorized, "Unauthorized", gin.H{"error": "Token expired"})
			c.Abort()
			return
		}

		c.Set("userID", jsonToken.Subject)
		c.Next()
	}
}
