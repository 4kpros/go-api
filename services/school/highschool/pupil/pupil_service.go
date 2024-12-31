package pupil

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/highschool/pupil/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Create new pupil
func (service *Service) Create(inputJwtToken *types.JwtToken, pupil *model.Pupil) (result *model.Pupil, errCode int, err error) {
	// Check if pupil already exists
	foundPupil, err := service.Repository.GetByObject(&model.Pupil{
		SchoolID: pupil.SchoolID,
		UserID:   pupil.UserID,
		ClassID:  pupil.ClassID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get pupil by name from database")
		return
	}
	if foundPupil != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("pupil")
		return
	}

	// Insert pupil
	result, err = service.Repository.Create(pupil)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create pupil from database")
		return
	}
	return
}

// Update pupil
func (service *Service) Update(inputJwtToken *types.JwtToken, pupilID int64, pupil *model.Pupil) (result *model.Pupil, errCode int, err error) {
	// Check if pupil already exists
	foundPupilByID, err := service.Repository.GetById(pupilID, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get pupil by name from database")
		return
	}
	if foundPupilByID == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Pupil")
		return
	}
	foundPupil, err := service.Repository.GetByObject(&model.Pupil{
		SchoolID: foundPupilByID.SchoolID,
		UserID:   foundPupilByID.UserID,
		ClassID:  pupil.ClassID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get pupil by name from database")
		return
	}
	if foundPupil != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("pupil")
		return
	}

	// Update pupil
	result, err = service.Repository.Update(pupilID, inputJwtToken.UserID, pupil)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update pupil from database")
		return
	}
	return
}

// Delete pupil with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, pupilID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(pupilID, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete pupil from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Pupil")
		return
	}
	return
}

// Get Returns pupil with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, pupilID int64) (pupil *model.Pupil, errCode int, err error) {
	pupil, err = service.Repository.GetById(pupilID, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get pupil by id from database")
		return
	}
	if pupil == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Pupil")
		return
	}
	return
}

// GetAll Returns all faculties with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (pupilList []model.Pupil, errCode int, err error) {
	pupilList, err = service.Repository.GetAll(filter, pagination, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get faculties from database")
	}
	return
}
