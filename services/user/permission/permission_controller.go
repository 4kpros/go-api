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

func (controller *Controller) UpdateByRoleID(
	ctx *context.Context,
	roleID int64,
	feature string,
	body data.UpdateRoleFeaturePermissionBodyRequest,
) (result *data.PermissionFeatureTableResponse, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateByRoleID(
		helpers.GetJwtContext(ctx),
		roleID,
		feature,
		body,
	)
	return
}

func (controller *Controller) GetAllByRoleID(
	ctx *context.Context,
	input *struct {
		data.GetRolePermissionListRequest
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
