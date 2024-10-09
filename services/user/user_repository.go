package user

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/services/user/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	UpdateUser(user *model.User) error
	UpdateUserInfo(userInfo *model.UserInfo) error
	Delete(id string) (int64, error)
	GetById(id string) (*model.User, error)
	GetUserInfoById(id string) (*model.UserInfo, error)
	GetByEmail(email string) (*model.User, error)
	GetByPhoneNumber(phoneNumber int) (*model.User, error)
	GetByProvider(provider string, providerUserId string) (*model.User, error)
	GetAll(filter *types.Filter, pagination *types.Pagination) ([]model.User, error)
}

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: db}
}

func (repository *UserRepositoryImpl) Create(user *model.User) error {
	return repository.Db.Create(user).Error
}

func (repository *UserRepositoryImpl) UpdateUser(user *model.User) error {
	return repository.Db.Model(user).Updates(user).Error
}

func (repository *UserRepositoryImpl) UpdateUserInfo(userInfo *model.UserInfo) error {
	return repository.Db.Model(userInfo).Updates(userInfo).Error
}

func (repository *UserRepositoryImpl) Delete(id string) (int64, error) {
	var user = &model.User{}
	var result = repository.Db.Where("id = ?", id).Delete(user)
	return result.RowsAffected, result.Error
}

func (repository *UserRepositoryImpl) GetById(id string) (*model.User, error) {
	var user = &model.User{}
	var result = repository.Db.Where("id = ?", id).Limit(1).Find(user)
	return user, result.Error
}

func (repository *UserRepositoryImpl) GetUserInfoById(id string) (*model.UserInfo, error) {
	var userInfo = &model.UserInfo{}
	var result = repository.Db.Where("id = ?", id).Limit(1).Find(userInfo)
	return userInfo, result.Error
}

func (repository *UserRepositoryImpl) GetByEmail(email string) (*model.User, error) {
	var user = &model.User{}
	var result = repository.Db.Where("email = ? AND (provider is null OR provider = '')", email).Limit(1).Find(user)
	return user, result.Error
}

func (repository *UserRepositoryImpl) GetByPhoneNumber(phoneNumber int) (*model.User, error) {
	var user = &model.User{}
	var result = repository.Db.Where("phoneNumber = ? AND (provider is null OR provider = '')", phoneNumber).Limit(1).Find(user)
	return user, result.Error
}

func (repository *UserRepositoryImpl) GetByProvider(provider string, providerUserId string) (*model.User, error) {
	var user = &model.User{}
	var result = repository.Db.Where("provider = ? AND providerUserId = ?", provider, providerUserId).Limit(1).Find(user)
	return user, result.Error
}

func (repository *UserRepositoryImpl) GetAll(filter *types.Filter, pagination *types.Pagination) ([]model.User, error) {
	var users = []model.User{}
	var result = repository.Db.Scopes(utils.PaginationScope(users, pagination, filter, repository.Db)).Find(users)
	return users, result.Error
}
