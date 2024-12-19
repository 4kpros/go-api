package department

import (
	"context"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/university/department/data"
	"api/services/school/university/department/model"
)

type Controller struct {
	Service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{Service: service}
}

func (controller *Controller) Create(
	ctx *context.Context,
	input *struct {
		Body data.CreateDepartmentRequest
	},
) (result *model.Department, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		helpers.GetJwtContext(ctx),
		&model.Department{
			FacultyID:   input.Body.FacultyID,
			Name:        input.Body.Name,
			Description: input.Body.Description,
		},
	)
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.DepartmentID
		Body data.UpdateDepartmentRequest
	},
) (result *model.Department, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		helpers.GetJwtContext(ctx), input.ID,
		&model.Department{
			FacultyID:   input.ID,
			Name:        input.Body.Name,
			Description: input.Body.Description,
		},
	)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.DepartmentID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.Delete(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		data.DepartmentID
	},
) (result *model.Department, errCode int, err error) {
	department, errCode, err := controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = department
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
	},
) (result *data.DepartmentResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	departmentList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination)
	if err != nil {
		return
	}
	result = &data.DepartmentResponseList{
		Data: model.ToResponseList(departmentList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
