package role

import (
	"api/common/constants"
	"api/common/helpers"
	"api/common/types"
	"api/services/admin/role/model"
	"gorm.io/gorm"
)

type RoleRepository struct {
	Db                  *gorm.DB
	PermissionTableName string
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{Db: db, PermissionTableName: constants.PERMISSION_TABLE_NAME_ROLE}
}

func (repository *RoleRepository) Create(role *model.Role) (*model.Role, error) {
	result := *role
	return &result, repository.Db.Create(&result).Error
}

func (repository *RoleRepository) Update(id int64, role *model.Role) (*model.Role, error) {
	result := &model.Role{}
	return result, repository.Db.Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			"name":        role.Name,
			"description": role.Description,
		},
	).Error
}

func (repository *RoleRepository) Delete(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.Role{})
	return result.RowsAffected, result.Error
}

func (repository *RoleRepository) GetById(id int64) (*model.Role, error) {
	result := &model.Role{}
	return result, repository.Db.Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *RoleRepository) GetByName(name string) (*model.Role, error) {
	result := &model.Role{}
	return result, repository.Db.Where("name = ?", name).Limit(1).Find(result).Error
}

func (repository *RoleRepository) GetAll(filter *types.Filter, pagination *types.Pagination) ([]model.Role, error) {
	result := []model.Role{}
	return result, repository.Db.Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).Find(result).Error
}
