package history

import (
	"api/services/admin/history/model"
	"net/http"

	"api/common/constants"
	"api/common/types"
)

type HistoryService struct {
	Repository *HistoryRepository
}

func NewHistoryService(repository *HistoryRepository) *HistoryService {
	return &HistoryService{Repository: repository}
}

// Return all history with support for search, filter and pagination
func (service *HistoryService) GetAll(jwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (result []model.History, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("get all history from database")
	}
	return
}
