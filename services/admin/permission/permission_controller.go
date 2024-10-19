package permission

import (
	"context"

	"api/common/helpers"
	"api/common/types"
	"api/services/admin/permission/data"
	"api/services/admin/permission/model"
)

type Controller struct {
	Service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{Service: service}
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		Body data.UpdatePermissionRequest
	},
) (result *model.Permission, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		helpers.GetJwtContext(ctx),
		&model.Permission{
			RoleId:           input.Body.RoleId,
			FeatureName:      input.Body.FeatureName,
			TablePermissions: input.Body.TablePermissions,
		},
	)
	return
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		data.PermissionId
	},
) (result *model.Permission, errCode int, err error) {
	result, errCode, err = controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
	return
}

func (controller *Controller) GetAll(
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
