package role

import (
	"strconv"

	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/services/role/data"
	"github.com/4kpros/go-api/services/role/model"
)

type RoleController struct {
	Service RoleService
}

func NewRoleController(service RoleService) *RoleController {
	return &RoleController{Service: service}
}

func (controller *RoleController) Create(input *data.RoleRequest) (result *model.Role, errCode int, err error) {
	// Extract inputs
	var role = model.Role{}
	role.Name = (*input).Name
	role.Description = (*input).Description

	errCode, err = controller.Service.Create(&role)
	if err != nil {
		return
	}
	result = &role

	return
}

func (controller *RoleController) Update(input *model.Role) (result *model.Role, errCode int, err error) {
	// Extract inputs
	var role = *input

	errCode, err = controller.Service.Update(&role)
	if err != nil {
		return
	}
	result = &role

	return
}

func (controller *RoleController) Delete(input *data.RoleId) (result int64, errCode int, err error) {
	rows, errCode, err := controller.Service.Delete(strconv.Itoa(input.Id))
	if err != nil {
		return
	}
	result = rows

	return
}

func (controller *RoleController) FindById(input *data.RoleId) (result *model.Role, errCode int, err error) {
	role, errCode, err := controller.Service.FindById(strconv.Itoa(input.Id))
	if err != nil {
		return
	}
	result = role

	return
}

func (controller *RoleController) FindAll(filter *types.Filter, pagination *types.PaginationRequest) (result *data.RolesResponse, errCode int, err error) {
	// Calculate pagination
	newPagination, NewFilter := utils.GetPaginationFiltersFromQuery(filter, pagination)

	roles, errCode, err := controller.Service.FindAll(NewFilter, newPagination)
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
