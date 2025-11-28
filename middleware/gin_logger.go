package middleware

import (
	"scaffold/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		cost := time.Since(start)
		logger.Infof("status:%d method:%s path:%s query:%s ip:%s cost:%v", c.Writer.Status(), c.Request.Method, path, query, c.ClientIP(), cost)
	}
}
