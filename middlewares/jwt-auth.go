package middlewares

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthorizeJWT validates the token from the http request, returning a 401 if it's not valid
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Signing method validation
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			// I have to return the signing key
			return []byte("secret"), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println("Claims[Name]: ", claims["name"])
			fmt.Println("Claims[Admin]: ", claims["admin"])
			fmt.Println("Claims[Issuer]: ", claims["iss"])
			fmt.Println("Claims[IssuedAt]: ", claims["iat"])
			fmt.Println("Claims[ExpiresAt]: ", claims["exp"])
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
