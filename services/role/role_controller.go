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

func (controller *RoleController) Create(input *data.RoleRequest) (response *model.Role, errCode int, err error) {
	// Extract inputs
	var role = model.Role{}
	role.Name = (*input).Name
	role.Description = (*input).Description

	// Execute the service
	errCode, err = controller.Service.Create(&role)
	if err != nil {
		return
	}
	response = &role
	return
}

func (controller *RoleController) Update(input *model.Role) (response *model.Role, errCode int, err error) {
	// Extract inputs
	var role = *input

	// Execute the service
	errCode, err = controller.Service.Update(&role)
	if err != nil {
		return
	}
	response = &role
	return
}

func (controller *RoleController) Delete(input *data.Id) (response int64, errCode int, err error) {
	rows, errCode, err := controller.Service.Delete(strconv.Itoa(input.Id))
	if err != nil {
		return
	}
	response = rows
	return
}

func (controller *RoleController) FindById(input *data.Id) (response *model.Role, errCode int, err error) {
	role, errCode, err := controller.Service.FindById(strconv.Itoa(input.Id))
	if err != nil {
		return
	}
	response = role
	return
}

func (controller *RoleController) FindAll(filter *types.Filter, pagination *types.PaginationRequest) (response *data.GetAllResponse, errCode int, err error) {
	// Calculate pagination
	newPagination, NewFilter := utils.GetPaginationFiltersFromQuery(filter, pagination)

	// Execute the service
	roles, errCode, err := controller.Service.FindAll(NewFilter, newPagination)
	if err != nil {
		return
	}
	response = &data.GetAllResponse{
		Data: roles,
	}
	response.Filter = NewFilter
	response.Pagination = newPagination
	return
}
