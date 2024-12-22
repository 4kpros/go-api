package school

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/common/utils"
	"api/services/school/common/school/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Create new school
func (service *Service) Create(inputJwtToken *types.JwtToken, school *model.School) (result *model.School, errCode int, err error) {
	// Create school
	result, err = service.Repository.Create(school)
	if err != nil {
		pgState, errPgState := utils.ExtractSQLState(err.Error())
		if errPgState == nil {
			if pgState == constants.PG_ERROR_UNIQUE_COLUMN {
				errCode = http.StatusFound
				err = constants.Http302ErrorMessage("school")
				return
			}
		}
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create school from database")
		return
	}

	// Create config
	newConfig, _, err := service.UpdateConfig(inputJwtToken, -1, school.Config)
	if err != nil {
		return
	}

	// Create info
	newInfo, errCode, err := service.UpdateInfo(inputJwtToken, -1, school.Info)
	if err != nil {
		return
	}

	// Update school
	result, err = service.Repository.UpdateConfigInfoIDs(result.ID, newConfig.ID, newInfo.ID)
	result.Config = newConfig
	result.Info = newInfo
	return
}

// Update school
func (service *Service) Update(inputJwtToken *types.JwtToken, schoolID int64, school *model.School) (result *model.School, errCode int, err error) {
	// Check if school exists
	foundSchool, err := service.Repository.GetByName(school.Name)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get school by name from database")
		return
	}
	if foundSchool != nil && foundSchool.Name == school.Name {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("school")
		return
	}

	// Update school
	result, err = service.Repository.Update(schoolID, school)
	if err != nil {
		pgState, errPgState := utils.ExtractSQLState(err.Error())
		if errPgState == nil {
			if pgState == constants.PG_ERROR_UNIQUE_COLUMN {
				errCode = http.StatusFound
				err = constants.Http302ErrorMessage("school")
				return
			}
		}
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update school from database")
		return
	}

	// Update config
	newConfig, _, err := service.UpdateConfig(inputJwtToken, result.SchoolInfoID, school.Config)
	if err != nil {
		return
	}
	result.Config = newConfig

	// Update info
	newInfo, errCode, err := service.UpdateInfo(inputJwtToken, result.SchoolConfigID, school.Info)
	if err != nil {
		return
	}
	result.Info = newInfo
	return
}

// Update info
func (service *Service) UpdateInfo(inputJwtToken *types.JwtToken, id int64, item *model.SchoolInfo) (result *model.SchoolInfo, errCode int, err error) {
	// Check if school info already exists
	foundItem, err := service.Repository.GetInfoByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get school info by id from database")
		return
	}
	if foundItem != nil && foundItem.ID == id {
		// Update
		result, err = service.Repository.UpdateInfo(id, item)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage("update school info from database")
			return
		}
		return
	}

	// Create
	result, err = service.Repository.CreateInfo(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create school info from database")
		return
	}
	return
}

// Update config
func (service *Service) UpdateConfig(inputJwtToken *types.JwtToken, id int64, item *model.SchoolConfig) (result *model.SchoolConfig, errCode int, err error) {
	// Check if school config already exists
	foundItem, err := service.Repository.GetConfigByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get school config by id from database")
		return
	}
	if foundItem != nil && foundItem.ID == id {
		// Update
		result, err = service.Repository.UpdateConfig(id, item)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage("update school config from database")
			return
		}
		return
	}

	// Create
	result, err = service.Repository.CreateConfig(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create school config from database")
		return
	}
	return
}

// Delete school with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, schoolID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(schoolID)
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

// Get Returns school with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, schoolID int64) (school *model.School, errCode int, err error) {
	school, err = service.Repository.GetByID(schoolID)
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
