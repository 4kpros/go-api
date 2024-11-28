package role

import (
	"gorm.io/gorm"

	"api/common/helpers"
	"api/common/types"
	"api/services/user/role/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(role *model.Role) (*model.Role, error) {
	result := *role
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Update(roleID int64, role *model.Role) (*model.Role, error) {
	result := &model.Role{}
	return result, repository.Db.Model(result).Where("id = ?", roleID).Updates(
		map[string]interface{}{
			"name":        role.Name,
			"description": role.Description,
		},
	).Error
}

func (repository *Repository) Delete(roleID int64) (int64, error) {
	result := repository.Db.Where("id = ?", roleID).Delete(&model.Role{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetById(roleID int64) (*model.Role, error) {
	result := &model.Role{}
	return result, repository.Db.Where("id = ?", roleID).Limit(1).Find(result).Error
}

func (repository *Repository) GetByName(name string) (*model.Role, error) {
	result := &model.Role{}
	return result, repository.Db.Where("name = ?", name).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination) ([]model.Role, error) {
	var result []model.Role
	return result, repository.Db.Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).Find(result).Error
}
