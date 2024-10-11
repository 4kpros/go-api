package permission

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/services/permission/data"
	"github.com/4kpros/go-api/services/permission/model"
)

type PermissionController struct {
	Service PermissionService
}

func NewPermissionController(service PermissionService) *PermissionController {
	return &PermissionController{Service: service}
}

func (controller *PermissionController) Create(input *data.CreatePermissionRequest) (result *model.Permission, errCode int, err error) {
	var permission = model.Permission{
		RoleId: (*input).RoleId,
		Table:  (*input).Table,
		Read:   (*input).Read,
		Create: (*input).Create,
		Update: (*input).Update,
		Delete: (*input).Update,
	}
	errCode, err = controller.Service.Create(&permission)
	if err != nil {
		return
	}
	result = &permission
	return
}

func (controller *PermissionController) Update(id int64, input *data.UpdatePermissionRequest) (result *model.Permission, errCode int, err error) {
	var permission = model.Permission{
		Read:   (*input).Read,
		Create: (*input).Create,
		Update: (*input).Update,
		Delete: (*input).Update,
	}
	errCode, err = controller.Service.Update(id, input)
	if err != nil {
		return
	}
	result = &permission
	return
}

func (controller *PermissionController) Delete(id int64) (result int64, errCode int, err error) {
	var affectedRows int64
	affectedRows, errCode, err = controller.Service.Delete(id)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *PermissionController) GetById(id int64) (result *model.Permission, errCode int, err error) {
	var permission *model.Permission
	permission, errCode, err = controller.Service.GetById(id)
	if err != nil {
		return
	}
	result = permission
	return
}

func (controller *PermissionController) GetAll(filter *types.Filter, pagination *types.PaginationRequest) (result *data.PermissionsResponse, errCode int, err error) {
	var newPagination, NewFilter = utils.GetPaginationFiltersFromQuery(filter, pagination)
	var permissions []model.Permission
	permissions, errCode, err = controller.Service.GetAll(NewFilter, newPagination)
	if err != nil {
		return
	}
	result = &data.PermissionsResponse{
		Data: permissions,
	}
	result.Filter = NewFilter
	result.Pagination = newPagination
	return
}