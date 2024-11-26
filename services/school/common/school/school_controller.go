package school

import (
	"context"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/common/school/data"
	"api/services/school/common/school/model"
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
		Body data.SchoolRequest
	},
) (result *model.School, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		helpers.GetJwtContext(ctx),
		&model.School{
			Name: input.Body.Name,
			Type: input.Body.Type,
		},
	)
	return
}

func (controller *Controller) AddDirector(
	ctx *context.Context,
	input *struct {
		Body data.DirectorRequest
	},
) (result *model.SchoolDirector, errCode int, err error) {
	result, errCode, err = controller.Service.AddDirector(helpers.GetJwtContext(ctx), input.Body.SchoolId, input.Body.UserId)
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.SchoolId
		Body data.SchoolRequest
	},
) (result *model.School, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		helpers.GetJwtContext(ctx), input.ID,
		&model.School{
			Name: input.Body.Name,
			Type: input.Body.Type,
		},
	)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.SchoolId
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.Delete(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) DeleteDirector(
	ctx *context.Context,
	input *struct {
		data.DirectorRequest
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteDirector(helpers.GetJwtContext(ctx), input.SchoolId, input.UserId)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		data.SchoolId
	},
) (result *model.School, errCode int, err error) {
	school, errCode, err := controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = school
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
	},
) (result *data.SchoolResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	schoolList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination)
	if err != nil {
		return
	}
	result = &data.SchoolResponseList{
		Data: model.ToSchoolResponseList(schoolList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
