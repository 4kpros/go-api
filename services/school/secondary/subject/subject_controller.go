package subject

import (
	"context"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/secondary/subject/data"
	"api/services/school/secondary/subject/model"
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
		Body data.CreateSubjectRequest
	},
) (result *model.Subject, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		helpers.GetJwtContext(ctx),
		&model.Subject{
			SchoolID:     input.Body.SchoolID,
			ClassID:      input.Body.ClassID,
			Name:         input.Body.Name,
			Description:  input.Body.Description,
			Coefficient:  input.Body.Coefficient,
			Program:      input.Body.Program,
			Requirements: input.Body.Requirements,
		},
	)
	return
}

func (controller *Controller) AddProfessor(
	ctx *context.Context,
	input *struct {
		data.SubjectID
		Body data.SubjectProfessorRequest
	},
) (result *model.SubjectProfessor, errCode int, err error) {
	result, errCode, err = controller.Service.AddProfessor(
		helpers.GetJwtContext(ctx),
		&model.SubjectProfessor{
			SubjectID: input.ID,
			UserID:    input.Body.UserID,
		},
	)
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.SubjectID
		Body data.UpdateSubjectRequest
	},
) (result *model.Subject, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		helpers.GetJwtContext(ctx), input.ID,
		&model.Subject{
			Name:         input.Body.Name,
			Description:  input.Body.Description,
			Coefficient:  input.Body.Coefficient,
			Program:      input.Body.Program,
			Requirements: input.Body.Requirements,
		},
	)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.SubjectID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.Delete(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) DeleteProfessor(
	ctx *context.Context,
	input *struct {
		data.SubjectID
		Body data.SubjectProfessorRequest
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteProfessor(helpers.GetJwtContext(ctx), input.ID, input.Body.UserID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		data.SubjectID
	},
) (result *model.Subject, errCode int, err error) {
	subject, errCode, err := controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = subject
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
	},
) (result *data.SubjectResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	subjectList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination)
	if err != nil {
		return
	}
	result = &data.SubjectResponseList{
		Data: model.ToSubjectResponseList(subjectList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
