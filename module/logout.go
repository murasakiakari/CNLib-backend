package module

import (
	"CNLib-backend/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(router *gin.RouterGroup) {
	router.GET("/", AuthToken, func(c *gin.Context) {
		token := c.Request.Header.Get("Token")
		utility.ExpiredJWT(token)
		c.JSON(http.StatusOK, gin.H{"Message": "Logout success"})
	})
}
