package role

import (
	"fmt"
	"net/http"

	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/services/role/model"
)

type RoleService interface {
	Create(role *model.Role) (errCode int, err error)
	Update(role *model.Role) (errCode int, err error)
	Delete(id string) (affectedRows int64, errCode int, err error)
	GetById(id string) (role *model.Role, errCode int, err error)
	GetAll(filter *types.Filter, pagination *types.Pagination) (roles []model.Role, errCode int, err error)
}

type RoleServiceImpl struct {
	Repository RoleRepository
}

func NewRoleServiceImpl(repository RoleRepository) RoleService {
	return &RoleServiceImpl{Repository: repository}
}

func (service *RoleServiceImpl) Create(role *model.Role) (errCode int, err error) {
	// Check if role already exists
	var foundRole, errFound = service.Repository.GetByName(role.Name)
	if errFound != nil {
		errCode = http.StatusInternalServerError
		err = errFound
		return
	}
	if foundRole != nil && foundRole.Name == role.Name {
		var message = "This role already exists! Please enter anther one."
		errCode = http.StatusFound
		err = fmt.Errorf("%s", message)
		return
	}

	// Create new role
	err = service.Repository.Create(role)
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}
	return
}

func (service *RoleServiceImpl) Update(role *model.Role) (errCode int, err error) {
	err = service.Repository.Update(role)
	if err != nil {
		errCode = http.StatusInternalServerError
	}
	return
}

func (service *RoleServiceImpl) Delete(id string) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(id)
	if err != nil {
		errCode = http.StatusInternalServerError
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		var message = "Could not delete role that doesn't exists! Please enter valid id."
		err = fmt.Errorf("%s", message)
		return
	}
	return
}

func (service *RoleServiceImpl) GetById(id string) (role *model.Role, errCode int, err error) {
	role, err = service.Repository.GetById(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}
	if role == nil {
		errCode = http.StatusNotFound
		var message = "Role not found! Please enter valid id."
		err = fmt.Errorf("%s", message)
	}
	return
}

func (service *RoleServiceImpl) GetAll(filter *types.Filter, pagination *types.Pagination) (roles []model.Role, errCode int, err error) {
	roles, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
	}
	return
}
