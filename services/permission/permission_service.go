package permission

import (
	"net/http"

	"github.com/4kpros/go-api/common/constants"
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/services/permission/data"
	"github.com/4kpros/go-api/services/permission/model"
)

type PermissionService struct {
	Repository *PermissionRepository
}

func NewPermissionService(repository *PermissionRepository) *PermissionService {
	return &PermissionService{Repository: repository}
}

// Create new permission
func (service *PermissionService) Create(permission *model.Permission) (result *model.Permission, errCode int, err error) {
	// Check if permission already exists(unique by group of "roleId" and "table")
	foundPermission, err := service.Repository.GetByRoleIdTable(permission.RoleId, permission.Table)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("get permission by roleId and table from database")
		return
	}
	if foundPermission != nil {
		errCode = http.StatusFound
		err = constants.HTTP_302_ERROR_MESSAGE("permission")
		return
	}

	// Insert permission
	result, err = service.Repository.Create(permission)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("create permission from database")
		return
	}
	return
}

// Update permission
func (service *PermissionService) Update(id int64, data *data.UpdatePermissionRequest) (result *model.Permission, errCode int, err error) {
	// Check if permission already exists(unique by group of "roleId" and "table")
	foundPermission, err := service.Repository.GetById(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("get permission by roleId and table from database")
		return
	}
	if foundPermission == nil {
		errCode = http.StatusNotFound
		err = constants.HTTP_404_ERROR_MESSAGE("Permission")
		return
	}

	// Update only necessary fields
	foundPermission.Read = data.Read
	foundPermission.Create = data.Create
	foundPermission.Update = data.Update
	foundPermission.Delete = data.Delete
	result, err = service.Repository.Update(id, foundPermission)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("update permission from database")
		return
	}
	return
}

// Delete permission with matching id and return affected rows
func (service *PermissionService) Delete(id int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("delete permission from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.HTTP_404_ERROR_MESSAGE("Permission")
		return
	}
	return
}

// Return permission with matching id
func (service *PermissionService) GetById(id int64) (result *model.Permission, errCode int, err error) {
	result, err = service.Repository.GetById(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("get permission by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.HTTP_404_ERROR_MESSAGE("Permission")
		return
	}
	return
}

// Return all permissions with support for search, filter and pagination
func (service *PermissionService) GetAll(filter *types.Filter, pagination *types.Pagination) (result []model.Permission, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("get permissions from database")
	}
	return
}
