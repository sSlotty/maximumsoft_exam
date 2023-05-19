package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"goenv/responses"
	"goenv/services"
	"net/http"
)

func AuthorizeJWT() gin.HandlerFunc {

	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")

		if len(authHeader) <= len(BearerSchema) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, responses.ErrorResponse{Status: http.StatusUnauthorized, Message: "Invalid Token"})
			return
		}
		tokenString := authHeader[len(BearerSchema):]
		token, _ := services.JWTAuthService().ValidateToken(tokenString)

		if token == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, responses.ErrorResponse{Status: http.StatusUnauthorized, Message: "Invalid Token"})
			return
		}
		if token.Valid {
			_ = token.Claims.(jwt.MapClaims)
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, responses.ErrorResponse{Status: http.StatusUnauthorized, Message: "Invalid Token"})
			return
		}
	}
}
