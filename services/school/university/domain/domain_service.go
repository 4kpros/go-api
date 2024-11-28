package domain

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/university/domain/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Create new domain
func (service *Service) Create(inputJwtToken *types.JwtToken, domain *model.Domain) (result *model.Domain, errCode int, err error) {
	// Check if domain already exists
	foundDomain, err := service.Repository.GetByObject(&model.Domain{
		SchoolID:     domain.SchoolID,
		DepartmentID: domain.DepartmentID,
		Name:         domain.Name,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get domain by name from database")
		return
	}
	if foundDomain != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("domain")
		return
	}

	// Insert domain
	result, err = service.Repository.Create(domain)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create domain from database")
		return
	}
	return
}

// Update domain
func (service *Service) Update(inputJwtToken *types.JwtToken, id int64, domain *model.Domain) (result *model.Domain, errCode int, err error) {
	// Check if domain already exists
	foundDomainByID, err := service.Repository.GetById(id, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get domain by name from database")
		return
	}
	if foundDomainByID == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Domain")
		return
	}
	foundDomain, err := service.Repository.GetByObject(&model.Domain{
		SchoolID:     foundDomainByID.SchoolID,
		DepartmentID: foundDomainByID.DepartmentID,
		Name:         domain.Name,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get domain by name from database")
		return
	}
	if foundDomain != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("domain")
		return
	}

	// Update domain
	result, err = service.Repository.Update(id, inputJwtToken.UserID, domain)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update domain from database")
		return
	}
	return
}

// Delete domain with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(id, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete domain from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Domain")
		return
	}
	return
}

// Get Returns domain with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, id int64) (domain *model.Domain, errCode int, err error) {
	domain, err = service.Repository.GetById(id, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get domain by id from database")
		return
	}
	if domain == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Domain")
		return
	}
	return
}

// GetAll Returns all faculties with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (domainList []model.Domain, errCode int, err error) {
	domainList, err = service.Repository.GetAll(filter, pagination, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get faculties from database")
	}
	return
}
