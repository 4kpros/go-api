package role

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/services/role/data"
	"github.com/4kpros/go-api/services/role/model"
)

type RoleController struct {
	Service *RoleService
}

func NewRoleController(service *RoleService) *RoleController {
	return &RoleController{Service: service}
}

func (controller *RoleController) Create(input *data.RoleRequest) (result *model.Role, errCode int, err error) {
	role := model.Role{
		Name:        (*input).Name,
		Description: (*input).Description,
	}
	errCode, err = controller.Service.Create(&role)
	if err != nil {
		return
	}
	result = &role
	return
}

func (controller *RoleController) Update(input *model.Role) (result *model.Role, errCode int, err error) {
	role := *input
	errCode, err = controller.Service.Update(&role)
	if err != nil {
		return
	}
	result = &role
	return
}

func (controller *RoleController) Delete(input *data.RoleId) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.Delete(input.Id)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *RoleController) GetById(input *data.RoleId) (result *model.Role, errCode int, err error) {
	role, errCode, err := controller.Service.GetById(input.Id)
	if err != nil {
		return
	}
	result = role
	return
}

func (controller *RoleController) GetAll(filter *types.Filter, pagination *types.PaginationRequest) (result *data.RolesResponse, errCode int, err error) {
	newPagination, NewFilter := utils.GetPaginationFiltersFromQuery(filter, pagination)
	roles, errCode, err := controller.Service.GetAll(NewFilter, newPagination)
	if err != nil {
		return
	}
	result = &data.RolesResponse{
		Data: roles,
	}
	result.Filter = NewFilter
	result.Pagination = newPagination
	return
}
