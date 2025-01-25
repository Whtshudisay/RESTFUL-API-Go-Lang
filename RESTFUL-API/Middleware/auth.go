package Middleware

import (
	"awesomeProject/RESTFUL-API/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Not Authorized! "})
		return
	}
	userid, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Not Authorized !"})
		return
	}

	context.Set("UserID", userid)
	context.Next()
}
