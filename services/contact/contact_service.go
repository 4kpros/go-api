package contact

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/contact/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Create new contact
func (service *Service) Create(inputJwtToken *types.JwtToken, contact *model.Contact) (result *model.Contact, errCode int, err error) {
	// Insert contact
	result, err = service.Repository.Create(contact)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create contact from database")
		return
	}
	return
}

// Get Returns contact with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, contactID int64) (contact *model.Contact, errCode int, err error) {
	contact, err = service.Repository.GetById(contactID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get contact by id from database")
		return
	}
	if contact == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Contact")
		return
	}
	return
}

// GetAll Returns all contacts with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (contactList []model.Contact, errCode int, err error) {
	contactList, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get contacts from database")
	}
	return
}
