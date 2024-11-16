package permission

import (
	"context"

	"api/common/helpers"
	"api/common/types"
	"api/services/user/permission/data"
)

type Controller struct {
	Service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{Service: service}
}

func (controller *Controller) UpdateByRoleIdFeatureName(
	ctx *context.Context,
	roleId int64,
	featureName string,
	body data.UpdateRoleFeaturePermissionBodyRequest,
) (result *data.PermissionFeatureTableResponse, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateByRoleIdFeatureName(
		helpers.GetJwtContext(ctx),
		roleId,
		featureName,
		body,
	)
	return
}

func (controller *Controller) GetAllByRoleId(
	ctx *context.Context,
	input *struct {
		data.GetRolePermissionListRequest
		types.Filter
		types.PaginationRequest
	},
) (result *data.PermissionListResponse, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	permissionList, errCode, err := controller.Service.GetAllByRoleId(helpers.GetJwtContext(ctx), input.RoleId, newFilter, newPagination)
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
