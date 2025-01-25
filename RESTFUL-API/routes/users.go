package routes

import (
	"awesomeProject/RESTFUL-API/Model"
	"awesomeProject/RESTFUL-API/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signup(context *gin.Context) {
	var user Model.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data."})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not Save User."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "SignUp Successful"})
}

func login(context *gin.Context) {
	var user Model.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data."})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"Message": "Could not authenticate User"})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not authenticate User"})
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Login Successful", "Token": token})
}
