package class

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/secondary/class/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Create new class
func (service *Service) Create(inputJwtToken *types.JwtToken, class *model.Class) (result *model.Class, errCode int, err error) {
	// Check if class already exists
	foundClass, err := service.Repository.GetByObject(&model.Class{
		SchoolID: class.SchoolID,
		Name:     class.Name,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get class by name from database")
		return
	}
	if foundClass != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("class")
		return
	}

	// Insert class
	result, err = service.Repository.Create(class)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create class from database")
		return
	}
	return
}

// Update class
func (service *Service) Update(inputJwtToken *types.JwtToken, classID int64, class *model.Class) (result *model.Class, errCode int, err error) {
	// Check if class already exists
	foundClassByID, err := service.Repository.GetById(classID, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get class by name from database")
		return
	}
	if foundClassByID == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Class")
		return
	}
	foundClass, err := service.Repository.GetByObject(&model.Class{
		SchoolID: foundClassByID.SchoolID,
		Name:     class.Name,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get class by name from database")
		return
	}
	if foundClass != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("class")
		return
	}

	// Update class
	result, err = service.Repository.Update(classID, inputJwtToken.UserID, class)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update class from database")
		return
	}
	return
}

// Delete class with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, classID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(classID, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete class from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Class")
		return
	}
	return
}

// Get Returns class with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, classID int64) (class *model.Class, errCode int, err error) {
	class, err = service.Repository.GetById(classID, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get class by id from database")
		return
	}
	if class == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Class")
		return
	}
	return
}

// GetAll Returns all faculties with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (classList []model.Class, errCode int, err error) {
	classList, err = service.Repository.GetAll(filter, pagination, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get faculties from database")
	}
	return
}
