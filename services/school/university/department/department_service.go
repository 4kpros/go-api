package department

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/school"
	"api/services/school/university/department/model"
)

type Service struct {
	Repository       *Repository
	SchoolRepository *school.Repository
}

func NewService(repository *Repository, SchoolRepository *school.Repository) *Service {
	return &Service{
		Repository:       repository,
		SchoolRepository: SchoolRepository,
	}
}

// Create new department
func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.UniversityDepartment) (result *model.UniversityDepartment, errCode int, err error) {
	// Check if the school type is university
	foundSchool, err := service.SchoolRepository.GetByID(item.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get school by id from database")
		return
	}
	if foundSchool.Type != constants.SCHOOL_TYPE_UNIVERSITY {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
	}

	// Check if the new one exists
	foundNewDepartment, err := service.Repository.GetBySchoolIDFacultyIDName(item.SchoolID, item.FacultyID, item.Name)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get department by user school ids from database")
		return
	}
	if foundNewDepartment != nil && foundNewDepartment.SchoolID == item.SchoolID && foundNewDepartment.FacultyID == item.FacultyID && foundNewDepartment.Name == item.Name {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("Department")
		return
	}

	// Insert
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create department from database")
		return
	}
	return
}

// Update department
func (service *Service) Update(inputJwtToken *types.JwtToken, id int64, item *model.UniversityDepartment) (result *model.UniversityDepartment, errCode int, err error) {
	// Check if the school type is university
	foundSchool, err := service.SchoolRepository.GetByID(item.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get school by id from database")
		return
	}
	if foundSchool.Type != constants.SCHOOL_TYPE_UNIVERSITY {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
	}

	// Check if department exists
	foundDepartment, err := service.Repository.GetById(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get department by name from database")
		return
	}
	if foundDepartment == nil || foundDepartment.ID != id {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Department")
		return
	}
	// Check if the school type is university
	if foundDepartment.School.Type != constants.SCHOOL_TYPE_UNIVERSITY {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
	}

	// Check if the new one exists
	foundNewDepartment, err := service.Repository.GetBySchoolIDFacultyIDName(item.SchoolID, item.FacultyID, item.Name)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get department by user school ids from database")
		return
	}
	if foundNewDepartment != nil && foundNewDepartment.SchoolID == item.SchoolID && foundNewDepartment.FacultyID == item.FacultyID && foundNewDepartment.Name == item.Name {
		if !(foundDepartment.SchoolID == foundNewDepartment.SchoolID && foundDepartment.FacultyID == foundNewDepartment.FacultyID && foundDepartment.Name == foundNewDepartment.Name) {
			errCode = http.StatusFound
			err = constants.Http302ErrorMessage("Faculty")
			return
		}
	}

	// Update department
	result, err = service.Repository.Update(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update department from database")
		return
	}
	return
}

// Delete department with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete department from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Department")
		return
	}
	return
}

// Delete Deletes selection
func (service *Service) DeleteMultiple(inputJwtToken *types.JwtToken, list []int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultiple(list)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete multiple department from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Department selection")
		return
	}
	return
}

// Get Returns department with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, id int64) (result *model.UniversityDepartment, errCode int, err error) {
	result, err = service.Repository.GetById(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get department by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Department")
		return
	}
	return
}

// GetAll Returns all departments with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.UniversityDepartment, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, schoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get departments from database")
	}
	return
}
