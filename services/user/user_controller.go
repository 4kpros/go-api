package user

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/services/user/data"
	"github.com/4kpros/go-api/services/user/model"
)

type UserController struct {
	Service *UserService
}

func NewUserController(service *UserService) *UserController {
	return &UserController{Service: service}
}

func (controller *UserController) CreateWithEmail(data *data.CreateUserWithEmailRequest) (result *model.User, errCode int, err error) {
	user := &model.User{
		Email:  data.Email,
		RoleId: data.RoleId,
	}
	result, errCode, err = controller.Service.Create(user)
	return
}

func (controller *UserController) CreateWithPhoneNumber(data *data.CreateUserWithPhoneNumberRequest) (result *model.User, errCode int, err error) {
	user := &model.User{
		PhoneNumber: data.PhoneNumber,
		RoleId:      data.RoleId,
	}
	result, errCode, err = controller.Service.Create(user)
	return
}

func (controller *UserController) UpdateUser(user *model.User) (result *model.User, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateUser(user)
	return
}

func (controller *UserController) Delete(id int64) (result int64, errCode int, err error) {
	result, errCode, err = controller.Service.Delete(id)
	return
}

func (controller *UserController) GetById(id int64) (result *model.User, errCode int, err error) {
	result, errCode, err = controller.Service.GetById(id)
	return
}

func (controller *UserController) GetAll(filter *types.Filter, pagination *types.PaginationRequest) (result *data.UserResponseList, errCode int, err error) {
	newPagination, NewFilter := utils.GetPaginationFiltersFromQuery(filter, pagination)
	userList, errCode, err := controller.Service.GetAll(NewFilter, newPagination)
	if err != nil {
		return
	}
	result = &data.UserResponseList{
		Data: model.ToResponseList(userList),
	}
	result.Filter = NewFilter
	result.Pagination = newPagination
	return
}
