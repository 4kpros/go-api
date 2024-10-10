package permission

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/services/permission/model"
	"gorm.io/gorm"
)

type PermissionRepository struct {
	Db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) *PermissionRepository {
	return &PermissionRepository{Db: db}
}

func (repository *PermissionRepository) Create(permission *model.Permission) error {
	return repository.Db.Create(permission).Error
}

func (repository *PermissionRepository) Update(permission *model.Permission) error {
	return repository.Db.Model(permission).Updates(permission).Error
}

func (repository *PermissionRepository) Delete(id int64) (int64, error) {
	var permission = &model.Permission{}
	var result = repository.Db.Where("id = ?", id).Delete(permission)
	return result.RowsAffected, result.Error
}

func (repository *PermissionRepository) GetById(id int64) (*model.Permission, error) {
	var permission = &model.Permission{}
	var result = repository.Db.Where("id = ?", id).Limit(1).Find(permission)
	return permission, result.Error
}

func (repository *PermissionRepository) GetByRoleIdTable(roleId int64, table string) (*model.Permission, error) {
	var permissionInfo = &model.Permission{}
	var result = repository.Db.Where("roleId = ? AND table = ?", roleId, table).Limit(1).Find(permissionInfo)
	return permissionInfo, result.Error
}

func (repository *PermissionRepository) GetAll(filter *types.Filter, pagination *types.Pagination) ([]model.Permission, error) {
	var permissions = []model.Permission{}
	var result = repository.Db.Scopes(utils.PaginationScope(permissions, pagination, filter, repository.Db)).Find(permissions)
	return permissions, result.Error
}
