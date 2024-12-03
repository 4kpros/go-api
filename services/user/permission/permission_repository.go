package permission

import (
	"errors"

	"gorm.io/gorm"

	"api/common/helpers"
	"api/common/types"
	"api/services/user/permission/data"
	"api/services/user/permission/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) UpdateByRoleID(
	roleID int64,
	feature string,
	table data.UpdatePermissionTableRequest,
) (*data.PermissionFeatureTableResponse, error) {
	var err error
	result := &data.PermissionFeatureTableResponse{}
	// Update permission feature
	if tmpErr := repository.Db.Model(&model.PermissionFeature{}).
		Where("role_id = ?", roleID).
		Update("feature", feature).Error; tmpErr != nil {
		if errors.Is(tmpErr, gorm.ErrRecordNotFound) {
			newPermissionFeature := &model.PermissionFeature{
				RoleID:  roleID,
				Feature: feature,
			}
			err = repository.Db.Create(&newPermissionFeature).Error
			if err != nil {
				result.PermissionFeatureResponse = newPermissionFeature.ToResponse()
			}
		} else {
			err = tmpErr
		}
	}
	// Update permission table
	if tmpErr := repository.Db.Model(&model.PermissionTable{}).
		Where("role_id = ?", roleID).
		Where("table_name = ?", table.TableName).
		Updates(
			map[string]interface{}{
				"create": table.Create,
				"read":   table.Read,
				"update": table.Update,
				"delete": table.Delete,
			},
		).Error; tmpErr != nil {
		if errors.Is(tmpErr, gorm.ErrRecordNotFound) {
			newPermissionTable := &model.PermissionTable{
				RoleID: roleID,
				Create: table.Create,
				Read:   table.Read,
				Update: table.Update,
				Delete: table.Delete,
			}
			err = repository.Db.Create(&newPermissionTable).Error
			if err != nil {
				result.PermissionTableResponse = newPermissionTable.ToResponse()
			}
		} else {
			err = tmpErr
		}
	}
	return result, err
}

func (repository *Repository) GetPermissionFeatureByRoleID(
	roleID int64,
	feature string,
) (*model.PermissionFeature, error) {
	result := &model.PermissionFeature{}
	return result, repository.Db.Where("role_id = ?", roleID).First(result).Error
}

func (repository *Repository) GetPermissionTableByRoleIDAndTableName(
	roleID int64,
	tableName string,
) (*model.PermissionTable, error) {
	result := &model.PermissionTable{}
	return result, repository.Db.Where("role_id = ?", roleID).Where("table_name = ?", tableName).First(result).Error
}

func (repository *Repository) GetAllByRoleID(
	roleID int64,
	filter *types.Filter,
	pagination *types.Pagination,
) ([]data.PermissionFeatureTableResponse, error) {
	result := make([]data.PermissionFeatureTableResponse, 0)
	return result, repository.Db.Model(&model.PermissionFeature{}).
		Select("permission_features.*, permission_tables.*").Where("role_id = ?", roleID).
		Joins("INNER JOIN permission_tables ON permission_features.role_id = permission_tables.role_id").
		Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).
		Find(result).Error
}
