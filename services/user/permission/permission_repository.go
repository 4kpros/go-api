package permission

import (
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

func (repository *Repository) CreatePermissionFeature(permission *model.PermissionFeature) (*model.PermissionFeature, error) {
	result := *permission
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) CreatePermissionTable(permission *model.PermissionTable) (*model.PermissionTable, error) {
	result := *permission
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) UpdatePermissionFeature(
	roleID int64,
	feature string,
) (*model.PermissionFeature, error) {
	foundPermission, err := repository.GetPermissionFeature(roleID, feature)
	if err != nil {
		return nil, err
	}
	if foundPermission != nil {
		result := foundPermission
		result.RoleID = roleID
		result.Feature = feature
		return result, repository.Db.Model(&model.PermissionFeature{}).
			Where("role_id = ?", roleID).
			Update("feature", feature).Error
	}
	result, err := repository.CreatePermissionFeature(&model.PermissionFeature{RoleID: roleID, Feature: feature})
	return result, err
}

func (repository *Repository) UpdatePermissionTable(
	roleID int64,
	tableName string,
	data *model.PermissionTable,
) (*model.PermissionTable, error) {
	foundPermission, err := repository.GetPermissionTable(roleID, tableName)
	if err != nil {
		return nil, err
	}
	if foundPermission != nil {
		result := foundPermission
		result.RoleID = roleID
		result.TableName = tableName
		result.Create = data.Create
		result.Read = data.Read
		result.Update = data.Update
		result.Delete = data.Delete
		return result, repository.Db.Model(&model.PermissionTable{}).
			Where("role_id = ?", roleID).
			Where("table_name = ?", tableName).
			Updates(
				map[string]interface{}{
					"create": result.Create,
					"read":   result.Read,
					"update": result.Update,
					"delete": result.Delete,
				},
			).Error
	}
	result, err := repository.CreatePermissionTable(
		&model.PermissionTable{
			RoleID: roleID, TableName: tableName, Create: data.Create, Read: data.Read, Update: data.Update, Delete: data.Delete,
		},
	)
	return result, err
}

func (repository *Repository) GetPermissionFeature(
	roleID int64,
	feature string,
) (*model.PermissionFeature, error) {
	result := &model.PermissionFeature{}
	return result, repository.Db.Where("role_id = ?", roleID).First(result).Error
}

func (repository *Repository) GetPermissionTable(
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
