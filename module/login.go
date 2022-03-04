package module

import (
	"CNLib-backend/controller"
	"CNLib-backend/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(router *gin.RouterGroup) {
	router.POST("/", func(c *gin.Context) {
		err := c.Request.ParseForm()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": "Error appear when parsing the form"})
			return
		}
		account := c.Request.Form.Get("account")
		if account == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": "Error appear when getting the value account from form"})
			return
		}
		password := c.Request.Form.Get("password")
		if password == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": "Error appear when getting the value password from form"})
			return
		}
		preferPassword, findAccount := controller.UserMap[account]
		if !findAccount {
			c.JSON(http.StatusUnauthorized, gin.H{"Message": "Account not find"})
			return
		}
		if !controller.VerifyPassword(password, preferPassword) {
			c.JSON(http.StatusUnauthorized, gin.H{"Message": "Wrong password"})
			return
		}
		token, err := utility.GetJWT(account)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": "Error appear when getting token"})
			return
		}
		c.Header("Token", token)
		c.JSON(http.StatusOK, gin.H{"Message": "Login Success"})
	})
}

func AuthToken(c *gin.Context) {
	token := c.Request.Header.Get("Token")
	if utility.ValidateJWT(token) {
		c.Next()
	} else {
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{"Message": "Token not valid"})
	}
}
