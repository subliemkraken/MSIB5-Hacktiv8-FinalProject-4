package middleware

import (
	"FinalProject4/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", "Token Not Found!")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Bearer tokentokentoken
		arrayToken := strings.Split(authHeader, " ")

		var tokenString string
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := NewService().ValidateToken(tokenString)

		if err != nil {
			response := helper.APIResponse("Unauthorized", "Token Invalid!")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", "Token invalid")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userRole := claim["role"].(string)

		if userRole != "admin" {
			response := helper.APIResponse("Unauthorized", "User is not admin")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", userRole)
	}
}
