package permission

import (
	"github.com/4kpros/go-api/common/constants"
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/services/permission/model"
	"gorm.io/gorm"
)

type PermissionRepository struct {
	Db                  *gorm.DB
	PermissionTableName string
}

func NewPermissionRepository(db *gorm.DB) *PermissionRepository {
	return &PermissionRepository{Db: db, PermissionTableName: constants.PERMISSION_TABLE_NAME_PERMISSION}
}

func (repository *PermissionRepository) Create(permission *model.Permission) (*model.Permission, error) {
	result := *permission
	return &result, repository.Db.Create(&result).Error
}

func (repository *PermissionRepository) Update(id int64, permission *model.Permission) (*model.Permission, error) {
	result := &model.Permission{}
	return result, repository.Db.Model(result).Where(
		"id = ?", id,
	).Where(
		"role_id = ?", permission.RoleId,
	).Where(
		"table = ?", permission.Table,
	).Updates(
		map[string]interface{}{
			"create": permission.Create,
			"read":   permission.Read,
			"update": permission.Update,
			"delete": permission.Delete,
		},
	).Error
}

func (repository *PermissionRepository) GetById(id int64) (*model.Permission, error) {
	result := &model.Permission{}
	return result, repository.Db.Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *PermissionRepository) GetByRoleIdTable(roleId int64, table string) (*model.Permission, error) {
	result := &model.Permission{}
	return result, repository.Db.Where(
		"role_id = ?", roleId,
	).Where(
		"table = ?", table,
	).Limit(1).Find(result).Error
}

func (repository *PermissionRepository) GetAll(filter *types.Filter, pagination *types.Pagination) ([]model.Permission, error) {
	result := []model.Permission{}
	return result, repository.Db.Scopes(utils.PaginationScope(result, pagination, filter, repository.Db)).Find(result).Error
}
