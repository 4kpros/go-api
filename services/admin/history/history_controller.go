package history

import (
	"api/services/admin/history/data"
	"api/services/admin/history/model"
	"context"

	"api/common/helpers"
	"api/common/types"
)

type HistoryController struct {
	Service *HistoryService
}

func NewHistoryController(service *HistoryService) *HistoryController {
	return &HistoryController{Service: service}
}

func (controller *HistoryController) GetAll(
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
