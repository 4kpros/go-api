package year

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/year/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Create new year
func (service *Service) Create(inputJwtToken *types.JwtToken, year *model.Year) (result *model.Year, errCode int, err error) {
	// Check if year already exists
	foundYear, err := service.Repository.GetByObject(year)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get year by name from database")
		return
	}
	if foundYear != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("year")
		return
	}

	// Insert year
	result, err = service.Repository.Create(year)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create year from database")
		return
	}
	return
}

// Update year
func (service *Service) Update(inputJwtToken *types.JwtToken, id int64, year *model.Year) (result *model.Year, errCode int, err error) {
	// Check if year already exists
	foundYear, err := service.Repository.GetByObject(year)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get year by name from database")
		return
	}
	if foundYear != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("year")
		return
	}

	// Update year
	result, err = service.Repository.Update(id, year)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update year from database")
		return
	}
	return
}

// Delete year with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete year from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Year")
		return
	}
	return
}

// Get Returns year with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, id int64) (year *model.Year, errCode int, err error) {
	year, err = service.Repository.GetById(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get year by id from database")
		return
	}
	if year == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Year")
		return
	}
	return
}

// GetAll Returns all years with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (yearList []model.Year, errCode int, err error) {
	yearList, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get years from database")
	}
	return
}