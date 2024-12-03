package permission

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/user/permission/data"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// UpdateByRoleID Update permission
func (service *Service) UpdateByRoleID(
	inputJwtToken *types.JwtToken,
	roleID int64,
	feature string,
	body data.UpdateRoleFeaturePermissionBodyRequest,
) (result *data.PermissionFeatureTableResponse, errCode int, err error) {
	result, err = service.Repository.UpdateByRoleID(
		roleID, feature, body.Table,
	)
	return
}

// GetAllByRoleID GetAll Returns all permissions with matching role id and
// support for search, filter and pagination
func (service *Service) GetAllByRoleID(
	inputJwtToken *types.JwtToken,
	roleID int64,
	filter *types.Filter,
	pagination *types.Pagination,
) (result []data.PermissionFeatureTableResponse, errCode int, err error) {
	result, err = service.Repository.GetAllByRoleID(
		roleID, filter, pagination,
	)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get permissions from database")
	}
	return
}
