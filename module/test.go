package module

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(router *gin.RouterGroup) {
	router.GET("/", AuthToken, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Message": "Test success"})
	})
}
