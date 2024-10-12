package role

import (
	"net/http"

	"github.com/4kpros/go-api/common/constants"
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/services/role/model"
)

type RoleService struct {
	Repository *RoleRepository
}

func NewRoleService(repository *RoleRepository) *RoleService {
	return &RoleService{Repository: repository}
}

// Create new role
func (service *RoleService) Create(role *model.Role) (errCode int, err error) {
	// Check if role already exists
	foundRole, err := service.Repository.GetByName(role.Name)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("get role by name from database")
		return
	}
	if foundRole != nil {
		errCode = http.StatusFound
		err = constants.HTTP_302_ERROR_MESSAGE("role")
		return
	}

	// Insert role
	err = service.Repository.Create(role)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("create role from database")
		return
	}
	return
}

// Update role
func (service *RoleService) Update(role *model.Role) (errCode int, err error) {
	// Check if role already exists
	foundRole, err := service.Repository.GetByName(role.Name)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("get role by name from database")
		return
	}
	if foundRole != nil {
		errCode = http.StatusFound
		err = constants.HTTP_302_ERROR_MESSAGE("role")
		return
	}

	// Update role
	err = service.Repository.Update(role)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("update role from database")
		return
	}
	return
}

// Delete role with matching id and return affected rows
func (service *RoleService) Delete(id int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("delete role from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.HTTP_404_ERROR_MESSAGE("Role")
		return
	}
	return
}

// Return role with matching id
func (service *RoleService) GetById(id int64) (role *model.Role, errCode int, err error) {
	role, err = service.Repository.GetById(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("get role by id from database")
		return
	}
	if role == nil {
		errCode = http.StatusNotFound
		err = constants.HTTP_404_ERROR_MESSAGE("Role")
		return
	}
	return
}

// Return all roles with support for search, filter and pagination
func (service *RoleService) GetAll(filter *types.Filter, pagination *types.Pagination) (roles []model.Role, errCode int, err error) {
	roles, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("get roles from database")
	}
	return
}
