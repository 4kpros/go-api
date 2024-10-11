package history

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/services/history/data"
	"github.com/4kpros/go-api/services/history/model"
)

type HistoryController struct {
	Service *HistoryService
}

func NewHistoryController(service *HistoryService) *HistoryController {
	return &HistoryController{Service: service}
}

func (controller *HistoryController) Create(input *model.History) (result *model.History, errCode int, err error) {
	var history = *input
	errCode, err = controller.Service.Create(&history)
	if err != nil {
		return
	}
	result = &history
	return
}

func (controller *HistoryController) GetAll(filter *types.Filter, pagination *types.PaginationRequest) (result *data.HistoriesResponse, errCode int, err error) {
	var newPagination, NewFilter = utils.GetPaginationFiltersFromQuery(filter, pagination)
	var histories []model.History
	histories, errCode, err = controller.Service.GetAll(NewFilter, newPagination)
	if err != nil {
		return
	}
	result = &data.HistoriesResponse{
		Data: histories,
	}
	result.Filter = NewFilter
	result.Pagination = newPagination
	return
}
