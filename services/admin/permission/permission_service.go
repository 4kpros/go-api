package permission

import (
	"api/services/admin/permission/data"
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

// UpdateByRoleIdFeatureName Update permission
func (service *Service) UpdateByRoleIdFeatureName(
	jwtToken *types.JwtToken,
	roleId int64,
	featureName string,
	body data.UpdateRoleFeaturePermissionBodyRequest,
) (result *model.Permission, errCode int, err error) {
	// TODO
	return
}

// GetByRoleIdFeatureName Get Returns permission with matching id
func (service *Service) GetByRoleIdFeatureName(
	jwtToken *types.JwtToken,
	roleId int64,
	featureName string,
) (result *model.Permission, errCode int, err error) {
	result, err = service.Repository.GetByRoleIdFeatureName(
		roleId, featureName,
	)
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

// GetAllByRoleId GetAll Returns all permissions with matching role id and
// support for search, filter and pagination
func (service *Service) GetAllByRoleId(
	jwtToken *types.JwtToken,
	roleId int64,
	filter *types.Filter,
	pagination *types.Pagination,
) (result []model.Permission, errCode int, err error) {
	result, err = service.Repository.GetAllByRoleId(
		roleId, filter, pagination,
	)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get permissions from database")
	}
	return
}
