package role

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/services/role/model"
	"gorm.io/gorm"
)

type RoleRepository struct {
	Db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{Db: db}
}

func (repository *RoleRepository) Create(role *model.Role) error {
	return repository.Db.Create(role).Error
}

func (repository *RoleRepository) Update(role *model.Role) error {
	return repository.Db.Model(role).Updates(role).Error
}

func (repository *RoleRepository) Delete(id int64) (int64, error) {
	role := &model.Role{}
	result := repository.Db.Where("id = ?", id).Delete(role)
	return result.RowsAffected, result.Error
}

func (repository *RoleRepository) GetById(id int64) (*model.Role, error) {
	role := &model.Role{}
	result := repository.Db.Where("id = ?", id).Limit(1).Find(role)
	return role, result.Error
}

func (repository *RoleRepository) GetByName(name string) (*model.Role, error) {
	role := &model.Role{}
	result := repository.Db.Where("name = ?", name).Limit(1).Find(role)
	return role, result.Error
}

func (repository *RoleRepository) GetAll(filter *types.Filter, pagination *types.Pagination) ([]model.Role, error) {
	roles := []model.Role{}
	result := repository.Db.Scopes(utils.PaginationScope(roles, pagination, filter, repository.Db)).Find(roles)
	return roles, result.Error
}
