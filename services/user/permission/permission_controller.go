package permission

import (
	"context"

	"api/common/helpers"
	"api/common/types"
	"api/services/user/permission/data"
	"api/services/user/permission/model"
)

type Controller struct {
	Service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{Service: service}
}

func (controller *Controller) UpdatePermission(
	ctx *context.Context,
	input *struct {
		data.PermissionPathRequest
		Body data.UpdatePermissionRequest
	},
) (result *data.PermissionResponse, errCode int, err error) {
	tmpResult, errCode, err := controller.Service.UpdatePermission(
		helpers.GetJwtContext(ctx),
		input.RoleID,
		input.Body.TableName,
		&model.Permission{
			Create: input.Body.Create,
			Read:   input.Body.Read,
			Update: input.Body.Update,
			Delete: input.Body.Delete,
		},
	)
	result = tmpResult.ToResponse()
	return
}

func (controller *Controller) GetAllByRoleID(
	ctx *context.Context,
	input *struct {
		data.PermissionPathRequest
		types.Filter
		types.PaginationRequest
	},
) (result *data.PermissionListResponse, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	permissionList, errCode, err := controller.Service.GetAllByRoleID(helpers.GetJwtContext(ctx), input.RoleID, newFilter, newPagination)
	if err != nil {
		return
	}
	result = &data.PermissionListResponse{
		Data: permissionList,
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
