package user

import (
	"context"

	"github.com/4kpros/go-api/common/helpers"
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

func (controller *UserController) CreateWithEmail(
	ctx *context.Context,
	input *struct {
		Body data.CreateUserWithEmailRequest
	},
) (result *model.User, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		helpers.GetJwtContext(ctx),
		&model.User{
			Email:  input.Body.Email,
			RoleId: input.Body.RoleId,
		},
	)
	return
}

func (controller *UserController) CreateWithPhoneNumber(
	ctx *context.Context,
	input *struct {
		Body data.CreateUserWithPhoneNumberRequest
	},
) (result *model.User, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		helpers.GetJwtContext(ctx),
		&model.User{
			PhoneNumber: input.Body.PhoneNumber,
			RoleId:      input.Body.RoleId,
		},
	)
	return
}

func (controller *UserController) UpdateUser(
	ctx *context.Context,
	input *struct {
		data.UserId
		Body data.UpdateUserRequest
	},
) (result *model.User, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateUser(
		helpers.GetJwtContext(ctx),
		&model.User{
			Email:       input.Body.Email,
			PhoneNumber: input.Body.PhoneNumber,
			RoleId:      input.Body.RoleId,
		},
	)
	return
}

func (controller *UserController) Delete(
	ctx *context.Context,
	input *struct {
		data.UserId
	},
) (result int64, errCode int, err error) {
	result, errCode, err = controller.Service.Delete(
		helpers.GetJwtContext(ctx),
		input.ID,
	)
	return
}

func (controller *UserController) Get(
	ctx *context.Context,
	input *struct {
		data.UserId
	},
) (result *model.User, errCode int, err error) {
	result, errCode, err = controller.Service.Get(
		helpers.GetJwtContext(ctx),
		input.ID,
	)
	return
}

func (controller *UserController) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
	},
) (result *data.UserResponseList, errCode int, err error) {
	newPagination, newFilter := utils.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	userList, errCode, err := controller.Service.GetAll(
		helpers.GetJwtContext(ctx),
		newFilter,
		newPagination,
	)
	if err != nil {
		return
	}
	result = &data.UserResponseList{
		Data: model.ToResponseList(userList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
