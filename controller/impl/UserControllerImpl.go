package ControllerImpl

import (
	"errors"
	"net/http"
	"perpustakaan/helpers"
	"perpustakaan/model/web/request"
	"perpustakaan/model/web/response"
	"perpustakaan/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *userController {
	return &userController{service}
}

func (s *userController) GetAllUser(c *gin.Context) {
	users, err := s.service.FindAllUser()
	if err != nil {
		c.Error(err)
		return
	}
	if len(users) == 0 {
		c.Error(errors.New("not Found"))
		return
	}
	res := helpers.ApiResponse("Data Found", http.StatusOK, "success", response.ResponseUsers(users))
	c.JSON(200, res)
}

func (s *userController) CreateUser(c *gin.Context) {
	var input request.UserRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		validate := helpers.ValidatorHandlerError(err)
		respond := helpers.ApiResponse("Data Create Succssfully", http.StatusUnprocessableEntity, "SuccessFully", validate)
		c.JSON(http.StatusUnprocessableEntity, respond)
		return
	}

	newUser, err := s.service.CreateUser(input)
	if err != nil {
		c.Error(err)
		return
	}

	respond := helpers.ApiResponse("Data Create Succssfully", http.StatusOK, "SuccessFully", newUser)
	c.JSON(http.StatusOK, respond)
}
func (s *userController) FindOneUser(c *gin.Context) {
	Params := c.Param("id")

	Id, err := strconv.Atoi(Params)

	if err != nil {
		c.Error(err)
		return
	}
	User, err := s.service.FindOneUser(Id)
	if err != nil {
		c.Error(err)
		return

	}
	if User.Id == 0 {
		c.Error(errors.New("not Found"))
		return
	}
	respond := helpers.ApiResponse("User Found", http.StatusOK, "Successfully", response.ResponseUser(User))
	c.JSON(http.StatusOK, respond)
}
func (s *userController) UpdateUser(c *gin.Context) {
	var input request.UserUpdateRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		validate := helpers.ValidatorHandlerError(err)
		respond := helpers.ApiResponse("Data Create Succssfully", http.StatusUnprocessableEntity, "SuccessFully", validate)
		c.JSON(http.StatusUnprocessableEntity, respond)
		return
	}
	Params := c.Param("id")
	Id, err := strconv.Atoi(Params)
	if err != nil {
		c.Error(err)
		return
	}
	User, err := s.service.UpdateUser(Id, input)
	if err != nil {
		c.Error(err)
		return
	}
	respond := helpers.ApiResponse("User Found", http.StatusOK, "Successfully", response.ResponseUser(User))
	c.JSON(http.StatusOK, respond)
}
func (s *userController) DeleteUser(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.Error(err)
		return
	}
	user, err := s.service.DeleteUser(id)
	if err != nil {
		c.Error(err)
		return
	}
	respond := helpers.ApiResponse("User Deleted", http.StatusOK, "Successfully", response.ResponseUser(user))
	c.JSON(http.StatusOK, respond)

}
