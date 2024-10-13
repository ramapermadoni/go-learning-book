package middlewares

import (
	"go-learning-book/utils/common"

	"github.com/gin-gonic/gin"
)

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		trace_id := common.GenerateRandomString(6)
		c.Set("trace_id", trace_id)

		c.Next()
	}
}
