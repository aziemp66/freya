package http

import (
	"github.com/gin-gonic/gin"

	errorCommon "github.com/aziemp66/freya-be/common/error"
)

func MiddlewareErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors[0]
			// if err can be casted to ClientError, then it is a client error
			if clientError, ok := err.Err.(errorCommon.ClientError); ok {
				c.JSON(clientError.Code, Error{
					Code:    clientError.Code,
					Message: clientError.Message,
				})
			} else if err.IsType(gin.ErrorTypeBind) {
				c.JSON(400, Error{
					Code:    400,
					Message: err.Err.Error(),
				})
			} else if err.IsType(gin.ErrorTypePrivate) {
				c.JSON(500, Error{
					Code:    500,
					Message: "Internal server error",
				})
			} else {
				c.JSON(500, Error{
					Code:    500,
					Message: "Internal server error",
				})
			}
		}
	}
}
