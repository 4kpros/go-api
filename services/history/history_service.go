package history

import (
	"net/http"

	"github.com/4kpros/go-api/common/constants"
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/services/history/model"
)

type HistoryService struct {
	Repository HistoryRepository
}

func NewHistoryService(repository HistoryRepository) *HistoryService {
	return &HistoryService{Repository: repository}
}

// Create new history entry
func (service *HistoryService) Create(history *model.History) (errCode int, err error) {
	err = service.Repository.Create(history)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("create history entry from database")
		return
	}
	return
}

// Return all history with support for search, filter and pagination
func (service *HistoryService) GetAll(filter *types.Filter, pagination *types.Pagination) (histories []model.History, errCode int, err error) {
	histories, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("get all history from database")
	}
	return
}
