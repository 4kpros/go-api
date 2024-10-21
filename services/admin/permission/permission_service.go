package permission

import (
	"api/services/admin/permission/data"
	"net/http"

	"api/common/constants"
	"api/common/types"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// UpdateByRoleIdFeatureName Update permission
func (service *Service) UpdateByRoleIdFeatureName(
	inputJwtToken *types.JwtToken,
	roleId int64,
	featureName string,
	body data.UpdateRoleFeaturePermissionBodyRequest,
) (result *data.PermissionFeatureTableResponse, errCode int, err error) {
	result, err = service.Repository.UpdateByRoleIdFeatureName(
		roleId, featureName, body.IsEnabled, body.Table,
	)
	return
}

// GetAllByRoleId GetAll Returns all permissions with matching role id and
// support for search, filter and pagination
func (service *Service) GetAllByRoleId(
	inputJwtToken *types.JwtToken,
	roleId int64,
	filter *types.Filter,
	pagination *types.Pagination,
) (result []data.PermissionFeatureTableResponse, errCode int, err error) {
	result, err = service.Repository.GetAllByRoleId(
		roleId, filter, pagination,
	)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get permissions from database")
	}
	return
}
