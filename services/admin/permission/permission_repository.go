package permission

import (
	"api/common/helpers"
	"api/services/admin/permission/data"
	"errors"
	"gorm.io/gorm"

	"api/common/types"
	"api/services/admin/permission/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) UpdateByRoleIdFeatureName(
	roleId int64,
	featureName string,
	isEnabled bool,
	table data.UpdatePermissionTableRequest,
) (*data.PermissionFeatureTableResponse, error) {
	var err error
	var result *data.PermissionFeatureTableResponse
	// Update permission feature
	if tmpErr := repository.Db.Model(&model.PermissionFeature{}).
		Where("role_id = ?", roleId).
		Where("feature_name = ?", featureName).
		Update("is_enabled", isEnabled).Error; tmpErr != nil {
		if errors.Is(tmpErr, gorm.ErrRecordNotFound) {
			newPermissionFeature := &model.PermissionFeature{
				RoleId:      roleId,
				FeatureName: featureName,
				IsEnabled:   isEnabled,
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
		Where("role_id = ?", roleId).
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
				RoleId: roleId,
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

func (repository *Repository) GetPermissionFeatureByRoleIdAndFeatureName(
	roleId int64,
	featureName string,
) (*model.PermissionFeature, error) {
	var result *model.PermissionFeature
	return result, repository.Db.Where("role_id = ?", roleId).Where("feature_name = ?", featureName).First(result).Error
}

func (repository *Repository) GetPermissionTableByRoleIdAndTableName(
	roleId int64,
	tableName string,
) (*model.PermissionTable, error) {
	var result *model.PermissionTable
	return result, repository.Db.Where("role_id = ?", roleId).Where("table_name = ?", tableName).First(result).Error
}

func (repository *Repository) GetAllByRoleId(
	roleId int64,
	filter *types.Filter,
	pagination *types.Pagination,
) ([]data.PermissionFeatureTableResponse, error) {
	var result []data.PermissionFeatureTableResponse
	return result, repository.Db.Model(&model.PermissionFeature{}).
		Select("permission_features.*, permission_tables.*").Where("role_id = ?", roleId).
		Joins("INNER JOIN permission_tables ON permission_features.role_id = permission_tables.role_id").
		Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).
		Find(result).Error
}
