package permission

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/admin/permission/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Update permission
func (service *Service) Update(jwtToken *types.JwtToken, permission *model.Permission) (result *model.Permission, errCode int, err error) {
	// Check if permission already exists(unique by group of "roleId" and "table")
	foundPermission, err := service.Repository.GetByRoleIdTable(permission.RoleId, permission.Table)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get permission by roleId and table from database")
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
			err = constants.Http500ErrorMessage("update permission from database")
			return
		}
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("permission")
		return
	}

	// Insert permission
	result, err = service.Repository.Create(permission)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create permission from database")
		return
	}
	return
}

// Get Returns permission with matching id
func (service *Service) Get(jwtToken *types.JwtToken, id int64) (result *model.Permission, errCode int, err error) {
	result, err = service.Repository.GetById(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get permission by id from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Permission")
		return
	}
	return
}

// GetAll Returns all permissions with support for search, filter and pagination
func (service *Service) GetAll(jwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (result []model.Permission, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get permissions from database")
	}
	return
}
