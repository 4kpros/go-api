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

func (repository *Repository) CreatePermission(permission *model.Permission) (*model.Permission, error) {
	result := *permission
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) UpdatePermission(
	roleID int64,
	tableName string,
	data *model.Permission,
) (*model.Permission, error) {
	result := &model.Permission{
		RoleID:    roleID,
		TableName: tableName,
		Create:    data.Create,
		Read:      data.Read,
		Update:    data.Update,
		Delete:    data.Delete,
	}
	return result, repository.Db.
		Where("role_id = ?", roleID).Where("table_name = ?", tableName).
		Attrs(
			model.Permission{
				Create: data.Create,
				Read:   data.Read,
				Update: data.Update,
				Delete: data.Delete,
			},
		).FirstOrCreate(result).Error
}

func (repository *Repository) GetPermission(
	roleID int64,
	tableName string,
) (*model.Permission, error) {
	result := &model.Permission{}
	return result, repository.Db.Where("role_id = ?", roleID).Where("table_name = ?", tableName).Limit(1).Find(result).Error
}

func (repository *Repository) GetPermissionOR(
	roleID int64,
	tableName1 string,
	tableName2 string,
) (*model.Permission, error) {
	result := &model.Permission{}
	return result, repository.Db.Where("role_id = ?", roleID).Where("table_name = ? or table_name = ?", tableName1, tableName2).Limit(1).Find(result).Error
}

func (repository *Repository) GetAllByRoleID(
	roleID int64,
	filter *types.Filter,
	pagination *types.Pagination,
) ([]data.PermissionResponse, error) {
	result := make([]data.PermissionResponse, 0)
	return result, repository.Db.Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).Where("role_id = ?", roleID).Find(result).Error
}
