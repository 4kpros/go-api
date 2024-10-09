package role

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/services/role/model"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(role *model.Role) error
	Update(role *model.Role) error
	Delete(id string) (int64, error)
	GetById(id string) (*model.Role, error)
	GetByName(name string) (*model.Role, error)
	GetAll(filter *types.Filter, pagination *types.Pagination) ([]model.Role, error)
}

type RoleRepositoryImpl struct {
	Db *gorm.DB
}

func NewRoleRepositoryImpl(db *gorm.DB) RoleRepository {
	return &RoleRepositoryImpl{Db: db}
}

func (repository *RoleRepositoryImpl) Create(role *model.Role) error {
	return repository.Db.Create(role).Error
}

func (repository *RoleRepositoryImpl) Update(role *model.Role) error {
	return repository.Db.Model(role).Updates(role).Error
}

func (repository *RoleRepositoryImpl) Delete(id string) (int64, error) {
	var role = &model.Role{}
	var result = repository.Db.Where("id = ?", id).Delete(role)
	return result.RowsAffected, result.Error
}

func (repository *RoleRepositoryImpl) GetById(id string) (*model.Role, error) {
	var role = &model.Role{}
	var result = repository.Db.Where("id = ?", id).Limit(1).Find(role)
	return role, result.Error
}

func (repository *RoleRepositoryImpl) GetByName(name string) (*model.Role, error) {
	var roleInfo = &model.Role{}
	var result = repository.Db.Where("name = ?", name).Limit(1).Find(roleInfo)
	return roleInfo, result.Error
}

func (repository *RoleRepositoryImpl) GetAll(filter *types.Filter, pagination *types.Pagination) ([]model.Role, error) {
	var roles = []model.Role{}
	var result = repository.Db.Scopes(utils.PaginationScope(roles, pagination, filter, repository.Db)).Find(roles)
	return roles, result.Error
}
