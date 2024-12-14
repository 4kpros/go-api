package permission

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/user/permission/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// UpdatePermission Update permission
func (service *Service) UpdatePermission(
	inputJwtToken *types.JwtToken,
	roleID int64,
	tableName string,
	data *model.Permission,
) (result *model.Permission, errCode int, err error) {
	result, err = service.Repository.UpdatePermission(
		roleID, tableName, data,
	)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update permissions from database")
	}
	return
}

// GetAll Returns all permissions with matching role id and
// support for search, filter and pagination
func (service *Service) GetAll(
	inputJwtToken *types.JwtToken,
	filter *types.Filter,
	pagination *types.Pagination,
) (result []model.Permission, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get permissions from database")
	}
	return
}
