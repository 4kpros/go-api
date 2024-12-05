package role

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/user/role/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Create new role
func (service *Service) Create(inputJwtToken *types.JwtToken, role *model.Role) (result *model.Role, errCode int, err error) {
	// Check if role already exists
	foundRole, err := service.Repository.GetByName(role.Name)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get role by name from database")
		return
	}
	if foundRole != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("role")
		return
	}

	// Insert role
	result, err = service.Repository.Create(role)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create role from database")
		return
	}
	return
}

// Update role
func (service *Service) Update(inputJwtToken *types.JwtToken, roleID int64, role *model.Role) (result *model.Role, errCode int, err error) {
	// Check if role already exists
	foundRole, err := service.Repository.GetByName(role.Name)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get role by name from database")
		return
	}
	if foundRole != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("role")
		return
	}

	// Update role
	result, err = service.Repository.Update(roleID, role)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update role from database")
		return
	}
	return
}

// Delete role with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, roleID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(roleID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete role from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Role")
		return
	}
	return
}

// Get Returns role with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, roleID int64) (role *model.Role, errCode int, err error) {
	role, err = service.Repository.GetById(roleID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get role by id from database")
		return
	}
	if role == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Role")
		return
	}
	return
}

// GetAll Returns all roles with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (roleList []model.Role, errCode int, err error) {
	roleList, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(err.Error())
		// err = constants.Http500ErrorMessage("get roles from database")
	}
	return
}
