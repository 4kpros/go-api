package faculty

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/university/faculty/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Create new faculty
func (service *Service) Create(inputJwtToken *types.JwtToken, faculty *model.Faculty) (result *model.Faculty, errCode int, err error) {
	// Check if faculty already exists
	foundFaculty, err := service.Repository.GetByObject(&model.Faculty{
		SchoolID: faculty.SchoolID,
		Name:     faculty.Name,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get faculty by name from database")
		return
	}
	if foundFaculty != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("faculty")
		return
	}

	// Insert faculty
	result, err = service.Repository.Create(faculty)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create faculty from database")
		return
	}
	return
}

// Update faculty
func (service *Service) Update(inputJwtToken *types.JwtToken, id int64, faculty *model.Faculty) (result *model.Faculty, errCode int, err error) {
	// Check if faculty already exists
	foundFacultyByID, err := service.Repository.GetById(id, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get faculty by name from database")
		return
	}
	if foundFacultyByID == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Faculty")
		return
	}
	foundFaculty, err := service.Repository.GetByObject(&model.Faculty{
		SchoolID: foundFacultyByID.SchoolID,
		Name:     faculty.Name,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get faculty by name from database")
		return
	}
	if foundFaculty != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("faculty")
		return
	}

	// Update faculty
	result, err = service.Repository.Update(id, inputJwtToken.UserID, faculty)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update faculty from database")
		return
	}
	return
}

// Delete faculty with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(id, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete faculty from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Faculty")
		return
	}
	return
}

// Get Returns faculty with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, id int64) (faculty *model.Faculty, errCode int, err error) {
	faculty, err = service.Repository.GetById(id, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get faculty by id from database")
		return
	}
	if faculty == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Faculty")
		return
	}
	return
}

// GetAll Returns all faculties with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (facultyList []model.Faculty, errCode int, err error) {
	facultyList, err = service.Repository.GetAll(filter, pagination, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get faculties from database")
	}
	return
}
