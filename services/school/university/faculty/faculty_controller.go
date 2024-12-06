package faculty

import (
	"context"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/university/faculty/data"
	"api/services/school/university/faculty/model"
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
		Body data.CreateFacultyRequest
	},
) (result *model.Faculty, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		helpers.GetJwtContext(ctx),
		&model.Faculty{
			SchoolID:    input.Body.SchoolID,
			Name:        input.Body.Name,
			Description: input.Body.Description,
		},
	)
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.FacultyID
		Body data.UpdateFacultyRequest
	},
) (result *model.Faculty, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		helpers.GetJwtContext(ctx), input.ID,
		&model.Faculty{
			Name:        input.Body.Name,
			Description: input.Body.Description,
		},
	)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.FacultyID
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
		data.FacultyID
	},
) (result *model.Faculty, errCode int, err error) {
	faculty, errCode, err := controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = faculty
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
	},
) (result *data.FacultyResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	facultyList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination)
	if err != nil {
		return
	}
	result = &data.FacultyResponseList{
		Data: model.ToResponseList(facultyList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
