package history

import (
	"net/http"

	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/services/history/model"
)

type HistoryService interface {
	Create(history *model.History) (errCode int, err error)
	GetAll(filter *types.Filter, pagination *types.Pagination) (histories []model.History, errCode int, err error)
}

type HistoryServiceImpl struct {
	Repository HistoryRepository
}

func NewHistoryServiceImpl(repository HistoryRepository) HistoryService {
	return &HistoryServiceImpl{Repository: repository}
}

func (service *HistoryServiceImpl) Create(history *model.History) (errCode int, err error) {
	err = service.Repository.Create(history)
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}
	return
}

func (service *HistoryServiceImpl) GetAll(filter *types.Filter, pagination *types.Pagination) (histories []model.History, errCode int, err error) {
	histories, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
	}
	return
}
