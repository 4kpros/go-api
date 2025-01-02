package section

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/school"
	"api/services/school/highschool/section/model"
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

// Create new section
func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.HighschoolSection) (result *model.HighschoolSection, errCode int, err error) {
	// Check if the school type is highschool
	foundSchool, err := service.SchoolRepository.GetByID(item.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get school by id from database")
		return
	}
	if foundSchool.Type != constants.SCHOOL_TYPE_HIGHSCHOOL {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
	}

	// Check if the new one exists
	foundNewSection, err := service.Repository.GetBySchoolIDName(item.SchoolID, item.Name)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get section by user school ids from database")
		return
	}
	if foundNewSection != nil && foundNewSection.SchoolID == item.SchoolID && foundNewSection.Name == item.Name {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("Section")
		return
	}

	// Insert
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create section from database")
		return
	}
	return
}

// Update section
func (service *Service) Update(inputJwtToken *types.JwtToken, id int64, item *model.HighschoolSection) (result *model.HighschoolSection, errCode int, err error) {
	// Check if the school type is highschool
	foundSchool, err := service.SchoolRepository.GetByID(item.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get school by id from database")
		return
	}
	if foundSchool.Type != constants.SCHOOL_TYPE_HIGHSCHOOL {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
	}

	// Check if section exists
	foundSection, err := service.Repository.GetById(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get section by name from database")
		return
	}
	if foundSection == nil || foundSection.ID != id {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Section")
		return
	}
	// Check if the school type is highschool
	if foundSection.School.Type != constants.SCHOOL_TYPE_HIGHSCHOOL {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
	}

	// Check if the new one exists
	foundNewSection, err := service.Repository.GetBySchoolIDName(item.SchoolID, item.Name)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get section by user school ids from database")
		return
	}
	if foundNewSection != nil && foundNewSection.SchoolID == item.SchoolID && foundNewSection.Name == item.Name {
		if !(foundSection.SchoolID == foundNewSection.SchoolID && foundSection.Name == foundNewSection.Name) {
			errCode = http.StatusFound
			err = constants.Http302ErrorMessage("Section")
			return
		}
	}

	// Update section
	result, err = service.Repository.Update(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update section from database")
		return
	}
	return
}

// Delete section with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete section from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Section")
		return
	}
	return
}

// Delete Deletes selection
func (service *Service) DeleteMultiple(inputJwtToken *types.JwtToken, list []int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultiple(list)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete multiple section from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Section selection")
		return
	}
	return
}

// Get Returns section with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, id int64) (result *model.HighschoolSection, errCode int, err error) {
	result, err = service.Repository.GetById(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get section by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Section")
		return
	}
	return
}

// GetAll Returns all sections with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.HighschoolSection, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, schoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get sections from database")
	}
	return
}
