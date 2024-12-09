package exam

import (
	"context"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/university/exam/data"
	"api/services/school/university/exam/model"
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
		Body data.CreateExamRequest
	},
) (result *model.Exam, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		helpers.GetJwtContext(ctx),
		&model.Exam{
			SchoolID:       input.Body.SchoolID,
			TeachingUnitID: input.Body.TeachingUnitID,
			Type:           input.Body.Type,
			Percentage:     input.Body.Percentage,
			Description:    input.Body.Description,
		},
	)
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.ExamID
		Body data.UpdateExamRequest
	},
) (result *model.Exam, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		helpers.GetJwtContext(ctx), input.ID,
		&model.Exam{
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
		data.ExamID
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
		data.ExamID
	},
) (result *model.Exam, errCode int, err error) {
	exam, errCode, err := controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = exam
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
	},
) (result *data.ExamResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	examList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination)
	if err != nil {
		return
	}
	result = &data.ExamResponseList{
		Data: model.ToResponseList(examList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
