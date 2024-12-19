package permission

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/user/permission/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Update Update permission
func (service *Service) Update(
	inputJwtToken *types.JwtToken,
	roleID int64,
	tableName string,
	data *model.Permission,
) (result *model.Permission, errCode int, err error) {
	// Check if the permission exists
	foundPermission, err := service.Repository.GetByRoleIDTableName(roleID, tableName)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("find permission from database")
		return
	}
	if foundPermission == nil || foundPermission.RoleID != roleID {
		// Create new ones
		result, err = service.Repository.Create(data)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage("create permission from database")
		}
		return
	}

	// Update now
	result, err = service.Repository.Update(
		roleID, tableName, data,
	)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update permission from database")
	}
	return
}

// Delete Deletes permission
func (service *Service) Delete(inputJwtToken *types.JwtToken, roleID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(roleID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete permission from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Role")
		return
	}
	return
}

// Delete Deletes selection
func (service *Service) DeleteMultiple(inputJwtToken *types.JwtToken, list []int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultiple(list)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete multiple permission from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Role selection")
		return
	}
	return
}

// GetAll Returns all permissions with matching role id and
// support for search, filter and pagination
func (service *Service) GetAll(
	inputJwtToken *types.JwtToken,
	filter *types.Filter,
	pagination *types.Pagination,
) (result []model.Permission, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get permissions from database")
	}
	return
}
