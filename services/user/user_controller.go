package user

import (
	"strconv"

	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/services/user/data"
	"github.com/4kpros/go-api/services/user/model"
)

type UserController struct {
	Service UserService
}

func NewUserController(service UserService) *UserController {
	return &UserController{Service: service}
}

func (controller *UserController) CreateWithEmail(input *data.UserWithEmailRequest) (result *model.User, errCode int, err error) {
	// Extract inputs
	var user = &model.User{}
	user.Email = input.Email
	user.Role = input.Role

	result, errCode, err = controller.Service.Create(user)

	return
}

func (controller *UserController) CreateWithPhoneNumber(input *data.UserWithPhoneNumberRequest) (result *model.User, errCode int, err error) {
	// Extract inputs
	var user = &model.User{}
	user.PhoneNumber = input.PhoneNumber
	user.Role = input.Role

	result, errCode, err = controller.Service.Create(user)

	return
}

func (controller *UserController) UpdateUser(input *model.User) (result *model.User, errCode int, err error) {
	// Extract inputs
	var user = *input

	errCode, err = controller.Service.UpdateUser(&user)
	if err != nil {
		return
	}
	result = &user

	return
}

func (controller *UserController) UpdateUserInfo(input *model.UserInfo) (result *model.UserInfo, errCode int, err error) {
	// Extract inputs
	var userInfo = *input

	errCode, err = controller.Service.UpdateUserInfo(&userInfo)
	if err != nil {
		return
	}
	result = &userInfo

	return
}

func (controller *UserController) Delete(input *data.UserId) (result int64, errCode int, err error) {
	result, errCode, err = controller.Service.Delete(strconv.Itoa(input.Id))

	return
}

func (controller *UserController) FindById(input *data.UserId) (result *model.User, errCode int, err error) {
	role, errCode, err := controller.Service.FindById(strconv.Itoa(input.Id))
	if err != nil {
		return
	}
	result = role

	return
}

func (controller *UserController) FindAll(filter *types.Filter, pagination *types.PaginationRequest) (result *data.UsersResponse, errCode int, err error) {
	// Calculate pagination
	newPagination, NewFilter := utils.GetPaginationFiltersFromQuery(filter, pagination)

	roles, errCode, err := controller.Service.FindAll(NewFilter, newPagination)
	if err != nil {
		return
	}
	result = &data.UsersResponse{
		Data: roles,
	}
	result.Filter = NewFilter
	result.Pagination = newPagination

	return
}
