package role

import (
	"context"

	"api/common/helpers"
	"api/common/types"
	"api/services/role/data"
	"api/services/role/model"
)

type RoleController struct {
	Service *RoleService
}

func NewRoleController(service *RoleService) *RoleController {
	return &RoleController{Service: service}
}

func (controller *RoleController) Create(
	ctx *context.Context,
	input *struct {
		Body data.RoleRequest
	},
) (result *model.Role, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		helpers.GetJwtContext(ctx),
		&model.Role{
			Name:        input.Body.Name,
			Description: input.Body.Description,
		},
	)
	return
}

func (controller *RoleController) Update(
	ctx *context.Context,
	input *struct {
		data.RoleId
		Body data.RoleRequest
	},
) (result *model.Role, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		helpers.GetJwtContext(ctx), input.ID,
		&model.Role{
			Name:        input.Body.Name,
			Description: input.Body.Description,
		},
	)
	return
}

func (controller *RoleController) Delete(
	ctx *context.Context,
	input *struct {
		data.RoleId
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.Delete(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *RoleController) Get(
	ctx *context.Context,
	input *struct {
		data.RoleId
	},
) (result *model.Role, errCode int, err error) {
	role, errCode, err := controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = role
	return
}

func (controller *RoleController) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
	},
) (result *data.RoleResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	roleList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination)
	if err != nil {
		return
	}
	result = &data.RoleResponseList{
		Data: model.ToResponseList(roleList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
