package permission

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/user/permission/data"
	"api/services/user/permission/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// UpdatePermissionFeature Update permission feature
func (service *Service) UpdatePermissionFeature(
	inputJwtToken *types.JwtToken,
	roleID int64,
	feature string,
) (result *model.PermissionFeature, errCode int, err error) {
	result, err = service.Repository.UpdatePermissionFeature(
		roleID, feature,
	)
	return
}

// UpdatePermissionTable Update permission table
func (service *Service) UpdatePermissionTable(
	inputJwtToken *types.JwtToken,
	roleID int64,
	tableName string,
	data *model.PermissionTable,
) (result *model.PermissionTable, errCode int, err error) {
	result, err = service.Repository.UpdatePermissionTable(
		roleID, tableName, data,
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
