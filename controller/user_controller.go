package controller

import (
	"FinalProject4/helper"
	"FinalProject4/middleware"
	"FinalProject4/model/input"
	"FinalProject4/model/response"
	"FinalProject4/service"
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{userService}
}

func (h *userController) RegisterUser(c *gin.Context) {
	var input input.UserRegisterInput

	err := c.ShouldBindJSON(&input)

	user, err := govalidator.ValidateStruct(input)

	if !user {
		c.JSON(http.StatusBadRequest, gin.H{
			"Errors": err.Error(),
		})
		fmt.Println("error: " + err.Error())
		return
	}

	result, err := h.userService.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Errors": err.Error(),
		})
		fmt.Println("error: " + err.Error())
		return
	}

	registerResponse := response.UserRegisterResponse{
		ID:        result.ID,
		FullName:  result.FullName,
		Email:     result.Email,
		Password:  result.Password,
		Balance:   result.Balance,
		CreatedAt: result.CreatedAt,
	}

	response := helper.APIResponse("created", registerResponse)
	c.JSON(201, response)
}

func (h *userController) LoginUser(c *gin.Context) {
	var input input.UserLoginInput

	err := c.ShouldBindJSON(&input)

	login, err := govalidator.ValidateStruct(input)

	if !login {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})

		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	// send to services
	// get user by email
	user, err := h.userService.GetUserByEmail(input.Email)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// return when user not found!
	if user.ID == 0 {
		errorMessages := "User not found!"
		response := helper.APIResponse("failed", errorMessages)
		c.JSON(http.StatusNotFound, response)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		response := helper.APIResponse("failed", "password not match!")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create token
	jwtService := middleware.NewService()
	token, err := jwtService.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		response := helper.APIResponse("failed", "failed to generate token!")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loginResponse := response.UserLoginResponse{
		Token: token,
	}

	// return token
	response := helper.APIResponse("ok", loginResponse)
	c.JSON(http.StatusOK, response)
}

func (h *userController) UserTopup(c *gin.Context) {
	var inputUserTopup input.UserTopupInput

	err := c.ShouldBindJSON(&inputUserTopup)

	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	user, err := govalidator.ValidateStruct(inputUserTopup)

	if !user {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	_, err = h.userService.UserTopup(currentUser, inputUserTopup)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})

		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	result, err := h.userService.GetUserByID(currentUser)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"Errors": err.Error(),
		})

		c.JSON(http.StatusBadRequest, response)
		fmt.Println("error: " + err.Error())
		return
	}

	userTopupResponse := response.UserTopupResponse{
		Message: "Top Up success! Your balance is: " + fmt.Sprintf("%d", result.Balance) + " now",
	}

	response := helper.APIResponse("ok", userTopupResponse)

	c.JSON(http.StatusOK, response)
}
