package handler

import (
	"crowfunding/helper"
	"crowfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)


type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors" : errors}

		response := helper.APIResponse("Registered Failed",http.StatusUnprocessableEntity,"Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Registered Failed",http.StatusBadRequest,"Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "tokentoken")

	response := helper.APIResponse("Account has been registered",http.StatusOK,"Success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors" : errors}

		response := helper.APIResponse("Login Failed",http.StatusUnprocessableEntity,"Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedInUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors" : err.Error()}

		response := helper.APIResponse("Login Failed",http.StatusUnprocessableEntity,"Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedInUser, "tokentoken")

	response := helper.APIResponse("User logged in",http.StatusOK,"Success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) CheckEmailAvailable(c *gin.Context)  {
	var input user.CheckEmailInput	

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors" : errors}

		response := helper.APIResponse("Email Checking Failed",http.StatusUnprocessableEntity,"Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors" : "Server Error"}

		response := helper.APIResponse("Email Checking Failed",http.StatusUnprocessableEntity,"Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}


	data := gin.H{
		"is_available" : isEmailAvailable,
	}

	
	metaMessage := "Email has been registered"
	
	if isEmailAvailable {
		metaMessage = "Email is available"
	} 
		
	response := helper.APIResponse(metaMessage ,http.StatusOK,"Success", data)
	c.JSON(http.StatusOK, response)
}