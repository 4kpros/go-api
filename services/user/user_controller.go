package user

import (
	"net/http"

	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/services/user/data/request"
	"github.com/4kpros/go-api/services/user/data/response"
	"github.com/4kpros/go-api/services/user/model"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service UserService
}

func NewUserController(service UserService) *UserController {
	return &UserController{Service: service}
}

// @Tags Users
// @Summary Create new user with email - [super-admin]
// @Accept  json
// @Produce  json
// @Param   payload body request.CreateWithEmailRequest true "Enter your information"
// @Success 200 {object} response.CreateWithEmailResponse "OK"
// @Failure 400 {object} types.ErrorResponse "Invalid email or role!"
// @Failure 302 {object} types.ErrorResponse "User with this email already exists!"
// @Security ApiKey && Bearer
// @Router /users/email [post]
func (controller *UserController) CreateWithEmail(c *gin.Context) {
	// Get data of req body
	var reqData = &request.CreateWithEmailRequest{}
	c.Bind(reqData)
	var user = &model.User{}
	user.Email = reqData.Email
	user.Role = reqData.Role

	// Execute the service
	password, errCode, err := controller.Service.CreateWithEmail(user)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, response.CreateWithEmailResponse{
		Email:    reqData.Email,
		Role:     reqData.Role,
		Password: password,
	})
}

// @Tags Users
// @Summary Create new user with phone number - [super-admin]
// @Accept  json
// @Produce  json
// @Param   payload body request.CreateWithPhoneNumberRequest true "Enter your information"
// @Success 200 {object} response.CreateWithPhoneNumberResponse "OK"
// @Failure 400 {object} types.ErrorResponse "Invalid phone number or role!"
// @Failure 302 {object} types.ErrorResponse "User with this phone number already exists!"
// @Security ApiKey && Bearer
// @Router /users/phone [post]
func (controller *UserController) CreateWithPhoneNumber(c *gin.Context) {
	// Get data of req body
	var reqData = &request.CreateWithPhoneNumberRequest{}
	c.Bind(reqData)
	var user = &model.User{}
	user.PhoneNumber = reqData.PhoneNumber
	user.Role = reqData.Role

	// Execute the service
	password, errCode, err := controller.Service.CreateWithPhoneNumber(user)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, response.CreateWithPhoneNumberResponse{
		PhoneNumber: reqData.PhoneNumber,
		Role:        reqData.Role,
		Password:    password,
	})
}

// @Tags Users
// @Summary Update user
// @Accept  json
// @Produce  json
// @Param   payload body model.User true "Enter your information"
// @Success 200 {object} model.User "OK"
// @Failure 400 {object} types.ErrorResponse "Invalid inputs!"
// @Failure 401 {object} types.ErrorResponse "Invalid user session!"
// @Failure 403 {object} types.ErrorResponse "Not permitted!"
// @Failure 404 {object} types.ErrorResponse "User not found!"
// @Security ApiKey && Bearer
// @Router /users/{id} [put]
func (controller *UserController) UpdateUser(c *gin.Context) {
	// Get data of req body
	var user = &model.User{}
	c.Bind(user)

	// Execute the service
	errCode, err := controller.Service.UpdateUser(user)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, user)
}

// @Tags Users
// @Summary Update user info
// @Accept  json
// @Produce  json
// @Param   payload body model.UserInfo false "User info model"
// @Success 200 {object} model.UserInfo "OK"
// @Failure 400 {object} types.ErrorResponse "Invalid inputs!"
// @Failure 401 {object} types.ErrorResponse "Invalid user session!"
// @Failure 403 {object} types.ErrorResponse "Not permitted!"
// @Failure 404 {object} types.ErrorResponse "User not found!"
// @Security ApiKey && Bearer
// @Router /users/info/{id} [put]
func (controller *UserController) UpdateUserInfo(c *gin.Context) {
	// Get data of req body
	var userInfo = &model.UserInfo{}
	c.Bind(userInfo)

	// Execute the service
	errCode, err := controller.Service.UpdateUserInfo(userInfo)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, userInfo)
}

// @Tags Users
// @Summary Update user info
// @Accept  json
// @Produce  json
// @Param   id path string true "User id"
// @Success 200 {object} types.ErrorResponse "OK"
// @Failure 400 {object} types.ErrorResponse "Invalid inputs!"
// @Failure 401 {object} types.ErrorResponse "Invalid user session!"
// @Failure 403 {object} types.ErrorResponse "Not permitted!"
// @Failure 404 {object} types.ErrorResponse "User not found!"
// @Security ApiKey && Bearer
// @Router /users/{id} [delete]
func (controller *UserController) Delete(c *gin.Context) {
	// Get data of req header
	id := c.Param("id")

	// Execute the service
	user, errCode, err := controller.Service.Delete(id)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, user)
}

// @Tags Users
// @Summary Get basic user info
// @Accept  json
// @Produce  json
// @Param   id path string true "User id"
// @Success 200 {object} model.User "OK"
// @Failure 400 {object} types.ErrorResponse "Invalid inputs!"
// @Failure 401 {object} types.ErrorResponse "Invalid user session!"
// @Failure 403 {object} types.ErrorResponse "Not permitted!"
// @Failure 404 {object} types.ErrorResponse "User not found!"
// @Security ApiKey && Bearer
// @Router /users/{id} [get]
func (controller *UserController) FindById(c *gin.Context) {
	// Get data of req header
	id := c.Param("id")

	// Execute the service
	user, errCode, err := controller.Service.FindById(id)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, user)
}

// @Tags Users
// @Summary Get detailed user info
// @Accept  json
// @Produce  json
// @Param   id path string true "UserInfo id"
// @Success 200 {object} model.UserInfo "OK"
// @Failure 400 {object} types.ErrorResponse "Invalid inputs!"
// @Failure 401 {object} types.ErrorResponse "Invalid user session!"
// @Failure 403 {object} types.ErrorResponse "Not permitted!"
// @Failure 404 {object} types.ErrorResponse "UserInfo not found!"
// @Security ApiKey && Bearer
// @Router /users/info/{id} [get]
func (controller *UserController) FindUserInfoById(c *gin.Context) {
	// Get data of req header
	id := c.Param("id")

	// Execute the service
	user, errCode, err := controller.Service.FindUserInfoById(id)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, user)
}

// @Tags Users
// @Summary Get all users
// @Accept  json
// @Produce  json
// @Param search query string false "Search keyword"
// @Param page query int false "Page"
// @Param limit query int false "Page limit"
// @Param orderBy query string false "Filter order by"
// @Param sort query string false "Sort asc, desc"
// @Success 200 {object} types.SuccessPaginatedResponse "OK"
// @Failure 401 {object} types.ErrorResponse "Invalid user session!"
// @Failure 403 {object} types.ErrorResponse "Not permitted!"
// @Security ApiKey && Bearer
// @Router /users/ [get]
func (controller *UserController) FindAll(c *gin.Context) {
	// Get data of req body
	pagination, filter := utils.GetPaginationFiltersFromQuery(c)

	// Execute the service
	users, errCode, err := controller.Service.FindAll(filter, pagination)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, types.SuccessPaginatedResponse{
		Data:       users,
		Filter:     filter,
		Pagination: pagination,
	})
}
