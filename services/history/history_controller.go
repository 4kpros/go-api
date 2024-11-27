package history

import (
	"context"

	"api/common/helpers"
	"api/common/types"
	"api/services/history/data"
	"api/services/history/model"
)

type Controller struct {
	Service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{Service: service}
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
	},
) (result *data.HistoryList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	historyList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination)
	if err != nil {
		return
	}
	result = &data.HistoryList{
		Data: model.ToResponseList(historyList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
