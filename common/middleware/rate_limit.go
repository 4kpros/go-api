package middleware

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/gin-gonic/gin"
)

func RateLimit(enabled bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !enabled {
			c.Next()
		}

		for _, err := range c.Errors {
			if err.Err.Error() != "EOF" {
				c.AbortWithStatusJSON(c.Writer.Status(), types.ErrorResponse{
					Message: err.Err.Error(),
				})
			}
		}
	}
}
