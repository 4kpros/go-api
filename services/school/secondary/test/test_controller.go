package test

import (
	"context"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/secondary/test/data"
	"api/services/school/secondary/test/model"
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
		Body data.CreateTestRequest
	},
) (result *model.Test, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		helpers.GetJwtContext(ctx),
		&model.Test{
			SchoolID:    input.Body.SchoolID,
			SubjectID:   input.Body.SubjectID,
			Type:        input.Body.Type,
			Percentage:  input.Body.Percentage,
			Description: input.Body.Description,
		},
	)
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.TestID
		Body data.UpdateTestRequest
	},
) (result *model.Test, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		helpers.GetJwtContext(ctx), input.ID,
		&model.Test{
			Type:        input.Body.Type,
			Percentage:  input.Body.Percentage,
			Description: input.Body.Description,
		},
	)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.TestID
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
		data.TestID
	},
) (result *model.Test, errCode int, err error) {
	test, errCode, err := controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = test
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
	},
) (result *data.TestResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	testList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination)
	if err != nil {
		return
	}
	result = &data.TestResponseList{
		Data: model.ToResponseList(testList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
