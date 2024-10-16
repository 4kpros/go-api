package permission

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/permission/model"
)

type PermissionService struct {
	Repository *PermissionRepository
}

func NewPermissionService(repository *PermissionRepository) *PermissionService {
	return &PermissionService{Repository: repository}
}

// Update permission
func (service *PermissionService) Update(jwtToken *types.JwtToken, permission *model.Permission) (result *model.Permission, errCode int, err error) {
	// Check if permission already exists(unique by group of "roleId" and "table")
	foundPermission, err := service.Repository.GetByRoleIdTable(permission.RoleId, permission.Table)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("get permission by roleId and table from database")
		return
	}
	if foundPermission != nil {
		// Update only necessary fields
		foundPermission.Read = permission.Read
		foundPermission.Create = permission.Create
		foundPermission.Update = permission.Update
		foundPermission.Delete = permission.Delete
		result, err = service.Repository.Update(foundPermission.ID, foundPermission)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.HTTP_500_ERROR_MESSAGE("update permission from database")
			return
		}
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

// Return permission with matching id
func (service *PermissionService) Get(jwtToken *types.JwtToken, id int64) (result *model.Permission, errCode int, err error) {
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
func (service *PermissionService) GetAll(jwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (result []model.Permission, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("get permissions from database")
	}
	return
}
