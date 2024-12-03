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

func (controller *Controller) UpdatePermissionFeature(
	ctx *context.Context,
	input *struct {
		data.PermissionPathRequest
		Body data.UpdatePermissionFeatureRequest
	},
) (result *data.PermissionFeatureResponse, errCode int, err error) {
	tmpResult, errCode, err := controller.Service.UpdatePermissionFeature(
		helpers.GetJwtContext(ctx),
		input.RoleID,
		input.Body.Feature,
	)
	result = tmpResult.ToResponse()
	return
}

func (controller *Controller) UpdatePermissionTable(
	ctx *context.Context,
	input *struct {
		data.PermissionPathRequest
		Body data.UpdatePermissionTableRequest
	},
) (result *data.PermissionTableResponse, errCode int, err error) {
	tmpResult, errCode, err := controller.Service.UpdatePermissionTable(
		helpers.GetJwtContext(ctx),
		input.RoleID,
		input.Body.TableName,
		&model.PermissionTable{
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
