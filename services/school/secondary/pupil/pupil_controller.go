package pupil

import (
	"context"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/secondary/pupil/data"
	"api/services/school/secondary/pupil/model"
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
		Body data.CreatePupilRequest
	},
) (result *model.Pupil, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		helpers.GetJwtContext(ctx),
		&model.Pupil{
			SchoolID: input.Body.SchoolID,
			UserID:   input.Body.UserID,
			ClassID:  input.Body.ClassID,
		},
	)
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.PupilID
		Body data.UpdatePupilRequest
	},
) (result *model.Pupil, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		helpers.GetJwtContext(ctx), input.ID,
		&model.Pupil{
			UserID:  input.Body.UserID,
			ClassID: input.Body.ClassID,
		},
	)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.PupilID
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
		data.PupilID
	},
) (result *model.Pupil, errCode int, err error) {
	pupil, errCode, err := controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = pupil
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
	},
) (result *data.PupilResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	pupilList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination)
	if err != nil {
		return
	}
	result = &data.PupilResponseList{
		Data: model.ToResponseList(pupilList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
