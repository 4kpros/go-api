package history

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/admin/history/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// GetAll Returns all history with support for search, filter and pagination
func (service *Service) GetAll(jwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (result []model.History, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get all history from database")
	}
	return
}
