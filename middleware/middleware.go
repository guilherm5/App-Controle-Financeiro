package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func MiddlewareGO() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := godotenv.Load("./.env")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"erro ao carregar variavel de ambiente para middleware": err,
			})
			log.Println("erro ao carregar variavel de ambiente para middleware", err)
		}

		secret := os.Getenv("SECRET")

		authHeader := c.GetHeader("Authorization")

		token, err := jwt.Parse(authHeader, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"Token invalido ou expirado": err,
			})
			log.Println("Token invalido ou expirado", err)
			return
		}
		c.Next()
	}
}
