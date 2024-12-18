package communication

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/communication/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Create new communication
func (service *Service) Create(inputJwtToken *types.JwtToken, communication *model.Communication) (result *model.Communication, errCode int, err error) {
	// Insert communication
	result, err = service.Repository.Create(communication)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create communication from database")
		return
	}
	return
}

// Get Returns communication with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, communicationID int64) (communication *model.Communication, errCode int, err error) {
	communication, err = service.Repository.GetById(communicationID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get communication by id from database")
		return
	}
	if communication == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Communication")
		return
	}
	return
}

// GetAll Returns all communications with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (communicationList []model.Communication, errCode int, err error) {
	communicationList, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get communications from database")
	}
	return
}
