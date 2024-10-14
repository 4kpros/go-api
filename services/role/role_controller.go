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
	result, errCode, err = controller.Service.Create(&role)
	return
}

func (controller *RoleController) Update(input *model.Role) (result *model.Role, errCode int, err error) {
	result, errCode, err = controller.Service.Update(input)
	return
}

func (controller *RoleController) Delete(input *data.RoleId) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.Delete(input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *RoleController) GetById(input *data.RoleId) (result *model.Role, errCode int, err error) {
	role, errCode, err := controller.Service.GetById(input.ID)
	if err != nil {
		return
	}
	result = role
	return
}

func (controller *RoleController) GetAll(filter *types.Filter, pagination *types.PaginationRequest) (result *data.RoleResponseList, errCode int, err error) {
	newPagination, NewFilter := utils.GetPaginationFiltersFromQuery(filter, pagination)
	roleList, errCode, err := controller.Service.GetAll(NewFilter, newPagination)
	if err != nil {
		return
	}
	result = &data.RoleResponseList{
		Data: model.ToResponseList(roleList),
	}
	result.Filter = NewFilter
	result.Pagination = newPagination
	return
}
