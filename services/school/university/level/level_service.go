package level

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/university/level/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Create new level
func (service *Service) Create(inputJwtToken *types.JwtToken, level *model.Level) (result *model.Level, errCode int, err error) {
	// Check if level already exists
	foundLevel, err := service.Repository.GetByObject(&model.Level{
		SchoolID: level.SchoolID,
		Name:     level.Name,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get level by name from database")
		return
	}
	if foundLevel != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("level")
		return
	}

	// Insert level
	result, err = service.Repository.Create(level)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create level from database")
		return
	}
	return
}

// Update level
func (service *Service) Update(inputJwtToken *types.JwtToken, levelID int64, level *model.Level) (result *model.Level, errCode int, err error) {
	// Check if level already exists
	foundLevelByID, err := service.Repository.GetById(levelID, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get level by name from database")
		return
	}
	if foundLevelByID == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Level")
		return
	}
	foundLevel, err := service.Repository.GetByObject(&model.Level{
		SchoolID: foundLevelByID.SchoolID,
		Name:     level.Name,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get level by name from database")
		return
	}
	if foundLevel != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("level")
		return
	}

	// Update level
	result, err = service.Repository.Update(levelID, inputJwtToken.UserID, level)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update level from database")
		return
	}
	return
}

// Delete level with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, levelID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(levelID, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete level from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Level")
		return
	}
	return
}

// Get Returns level with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, levelID int64) (level *model.Level, errCode int, err error) {
	level, err = service.Repository.GetById(levelID, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get level by id from database")
		return
	}
	if level == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Level")
		return
	}
	return
}

// GetAll Returns all faculties with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (levelList []model.Level, errCode int, err error) {
	levelList, err = service.Repository.GetAll(filter, pagination, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get faculties from database")
	}
	return
}
