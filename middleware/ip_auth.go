package middleware

import (
	"fmt"

	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
)

// IPAuthMiddleware 校验客户端IP访问名单
func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isMatched := false
		for _, host := range lib.GetStringSliceConf("base.http.allow_ip") {
			if c.ClientIP() == host {
				isMatched = true
				break
			}
		}
		if !isMatched {
			ResponseError(c, InternalErrorCode, fmt.Errorf("Access denied, ip(%v) not in allow_list", c.ClientIP()))
			c.Abort()
			return
		}
		c.Next()
	}
}
