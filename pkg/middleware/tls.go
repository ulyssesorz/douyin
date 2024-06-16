package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

// http重定向到https
func TLSSupportMiddleware(host string) gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     host,
		})

		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			return
		}

		c.Next()
	}
}
