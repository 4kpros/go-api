package user

import (
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
	var user = &model.User{
		Email:  input.Email,
		RoleId: input.RoleId,
	}
	result, errCode, err = controller.Service.Create(user)
	return
}

func (controller *UserController) CreateWithPhoneNumber(input *data.UserWithPhoneNumberRequest) (result *model.User, errCode int, err error) {
	var user = &model.User{
		PhoneNumber: input.PhoneNumber,
		RoleId:      input.RoleId,
	}
	result, errCode, err = controller.Service.Create(user)
	return
}

func (controller *UserController) UpdateUser(input *model.User) (result *model.User, errCode int, err error) {
	var user = *input
	errCode, err = controller.Service.UpdateUser(&user)
	if err != nil {
		return
	}
	result = &user
	return
}

func (controller *UserController) UpdateUserInfo(input *model.UserInfo) (result *model.UserInfo, errCode int, err error) {
	var userInfo = *input
	errCode, err = controller.Service.UpdateUserInfo(&userInfo)
	if err != nil {
		return
	}
	result = &userInfo
	return
}

func (controller *UserController) Delete(input *data.UserId) (result int64, errCode int, err error) {
	result, errCode, err = controller.Service.Delete(input.Id)
	return
}

func (controller *UserController) GetById(input *data.UserId) (result *model.User, errCode int, err error) {
	var user *model.User
	user, errCode, err = controller.Service.GetById(input.Id)
	if err != nil {
		return
	}
	result = user
	return
}

func (controller *UserController) GetAll(filter *types.Filter, pagination *types.PaginationRequest) (result *data.UsersResponse, errCode int, err error) {
	var newPagination, NewFilter = utils.GetPaginationFiltersFromQuery(filter, pagination)
	var users []model.User
	users, errCode, err = controller.Service.GetAll(NewFilter, newPagination)
	if err != nil {
		return
	}
	result = &data.UsersResponse{
		Data: users,
	}
	result.Filter = NewFilter
	result.Pagination = newPagination
	return
}
