package school

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/admin/school/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Create new school
func (service *Service) Create(inputJwtToken *types.JwtToken, school *model.School) (result *model.School, errCode int, err error) {
	// Check if school already exists
	foundSchool, err := service.Repository.GetByName(school.Name)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get school by name from database")
		return
	}
	if foundSchool != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("school")
		return
	}

	// Insert school
	result, err = service.Repository.Create(school)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create school from database")
		return
	}
	return
}

// Add school director
func (service *Service) AddDirector(inputJwtToken *types.JwtToken, schoolId int64, userId int64) (result *model.SchoolDirector, errCode int, err error) {
	// Check if director already exists
	foundSchoolDirector, err := service.Repository.GetDirector(schoolId, userId)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get school director from database")
		return
	}
	if foundSchoolDirector != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("school director")
		return
	}

	// Add new director
	result, err = service.Repository.AddDirector(schoolId, userId)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create school from database")
		return
	}
	return
}

// Update school
func (service *Service) Update(inputJwtToken *types.JwtToken, id int64, school *model.School) (result *model.School, errCode int, err error) {
	// Check if school already exists
	foundSchool, err := service.Repository.GetByName(school.Name)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get school by name from database")
		return
	}
	if foundSchool != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("school")
		return
	}

	// Update school
	result, err = service.Repository.Update(id, school)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update school from database")
		return
	}
	return
}

// Delete school with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete school from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("School")
		return
	}
	return
}

// Delete school director
func (service *Service) DeleteDirector(inputJwtToken *types.JwtToken, schoolId int64, userId int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteDirector(schoolId, userId)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete school director from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("School director")
		return
	}
	return
}

// Get Returns school with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, id int64) (school *model.School, errCode int, err error) {
	school, err = service.Repository.GetById(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get school by id from database")
		return
	}
	if school == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("School")
		return
	}
	return
}

// GetAll Returns all schools with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (schoolList []model.School, errCode int, err error) {
	schoolList, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get schools from database")
	}
	return
}
