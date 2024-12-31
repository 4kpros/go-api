package communication

import (
	"context"

	"api/common/helpers"
	"api/common/types"
	"api/services/communication/data"
	"api/services/communication/model"
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
		Body data.CommunicationRequest
	},
) (result *model.Communication, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		helpers.GetJwtContext(ctx),
		&model.Communication{
			Subject:       input.Body.Subject,
			Message:       input.Body.Message,
			AudienceType:  input.Body.AudienceType,
			AudienceValue: input.Body.AudienceValue,
		},
	)
	return
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		data.CommunicationID
	},
) (result *model.Communication, errCode int, err error) {
	communication, errCode, err := controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = communication
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
	},
) (result *data.CommunicationResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	communicationList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination)
	if err != nil {
		return
	}
	result = &data.CommunicationResponseList{
		Data: model.ToResponseList(communicationList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
