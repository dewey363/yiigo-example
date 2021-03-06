package middlewares

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/iiinsomnia/yiigo"
)

// Recovery panic recover middleware
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				yiigo.Logger.Error(fmt.Sprintf("yiigo demo panic: %v", err), zap.ByteString("stack", debug.Stack()))

				c.JSON(http.StatusOK, gin.H{
					"success": false,
					"code":    5000,
					"msg":     "服务器错误，请稍后重试",
				})

				return
			}
		}()

		c.Next()
	}
}
