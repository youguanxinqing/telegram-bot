package httputil

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequireToken(token string, handle func(c *gin.Context)) func(c *gin.Context) {
	return func(c *gin.Context) {
		if c.Param("token") != token {
			c.JSON(http.StatusForbidden, gin.H{"msg": "failed to check token"})
		} else {
			handle(c)
		}
	}
}
