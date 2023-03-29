package middleware

import (
	errorCommon "github.com/aziemp66/freya-be/common/error"
	"github.com/aziemp66/freya-be/common/jwt"
	"github.com/gin-gonic/gin"
)

func JWTAuth(j *jwt.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			authHeader = c.Param("token")
		}
		if len(authHeader) <= BEARER {
			c.Error(errorCommon.NewInvariantError("authorization header not valid"))
			c.Abort()
			return
		}
		tokenString := authHeader[BEARER:]
		claims, err := j.VerifyAuthToken(tokenString)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		c.Set("user_id", claims.ID)
		c.Set("user_role", claims.Role)
		c.Next()
	}
}

func RoleAuth(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.GetString("user_role")

		if userRole != role {
			c.Error(errorCommon.NewInvariantError("user not authorized"))
			c.Abort()
			return
		}

		c.Next()
	}
}
