package permission

import (
	"gorm.io/gorm"

	"api/common/helpers"
	"api/common/types"
	"api/services/admin/permission/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(permission *model.Permission) (*model.Permission, error) {
	result := *permission
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Update(id int64, permission *model.Permission) (*model.Permission, error) {
	return nil, nil
	//result := &model.Permission{}
	//return result, repository.Db.Model(result).Where(
	//	"id = ?", id,
	//).Where(
	//	"role_id = ?", permission.RoleId,
	//).Where(
	//	"table = ?", permission.Table,
	//).Updates(
	//	map[string]interface{}{
	//		"create": permission.Create,
	//		"read":   permission.Read,
	//		"update": permission.Update,
	//		"delete": permission.Delete,
	//	},
	//).Error
}

func (repository *Repository) GetById(id int64) (*model.Permission, error) {
	result := &model.Permission{}
	return result, repository.Db.Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByRoleIdTable(roleId int64, table string) (*model.Permission, error) {
	result := &model.Permission{}
	return result, repository.Db.Where(
		"role_id = ?", roleId,
	).Where(
		"table = ?", table,
	).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination) ([]model.Permission, error) {
	var result []model.Permission
	return result, repository.Db.Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).Find(result).Error
}
