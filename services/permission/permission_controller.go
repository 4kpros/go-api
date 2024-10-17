package permission

import (
	"context"

	"api/common/helpers"
	"api/common/types"
	"api/services/permission/data"
	"api/services/permission/model"
)

type PermissionController struct {
	Service *PermissionService
}

func NewPermissionController(service *PermissionService) *PermissionController {
	return &PermissionController{Service: service}
}

func (controller *PermissionController) Update(
	ctx *context.Context,
	input *struct {
		Body data.UpdatePermissionRequest
	},
) (result *model.Permission, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		helpers.GetJwtContext(ctx),
		&model.Permission{
			RoleId: input.Body.RoleId,
			Table:  input.Body.Table,
			Create: input.Body.Create,
			Read:   input.Body.Read,
			Update: input.Body.Update,
			Delete: input.Body.Delete,
		},
	)
	return
}

func (controller *PermissionController) Get(
	ctx *context.Context,
	input *struct {
		data.PermissionId
	},
) (result *model.Permission, errCode int, err error) {
	result, errCode, err = controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
	return
}

func (controller *PermissionController) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
	},
) (result *data.PermissionList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	permissionList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination)
	if err != nil {
		return
	}
	result = &data.PermissionList{
		Data: model.ToResponseList(permissionList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
