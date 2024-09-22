package auth

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/services/user/model"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Create(user *model.User) error
	CreateUserInfo(userInfo *model.UserInfo) error
	Update(user *model.User) error
	UpdatePasswordById(id string, password string) (*model.User, error)
	Delete(id string) (int64, error)
	FindById(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByPhoneNumber(phoneNumber int) (*model.User, error)
	FindByProvider(provider string, providerUserId string) (*model.User, error)
	FindAll(filter *types.Filter, pagination *types.Pagination) ([]model.User, error)
}

type AuthRepositoryImpl struct {
	Db *gorm.DB
}

func NewAuthRepositoryImpl(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{Db: db}
}

func (repository *AuthRepositoryImpl) Create(user *model.User) error {
	result := repository.Db.Create(user)
	return result.Error
}

func (repository *AuthRepositoryImpl) CreateUserInfo(userInfo *model.UserInfo) error {
	result := repository.Db.Create(userInfo)
	return result.Error
}

func (repository *AuthRepositoryImpl) Update(user *model.User) error {
	result := repository.Db.Model(user).Updates(user)
	return result.Error
}

func (repository *AuthRepositoryImpl) UpdatePasswordById(id string, password string) (*model.User, error) {
	var user = &model.User{
		Password: password,
	}
	result := repository.Db.Model(user).Where("id = ?", id).Update("password", user)
	return user, result.Error
}

func (repository *AuthRepositoryImpl) Delete(id string) (int64, error) {
	var user = &model.User{}
	result := repository.Db.Where("id = ?", id).Delete(user)
	return result.RowsAffected, result.Error
}

func (repository *AuthRepositoryImpl) FindAll(filter *types.Filter, pagination *types.Pagination) ([]model.User, error) {
	var users = []model.User{}
	result := repository.Db.Scopes(utils.PaginationScope(users, pagination, filter, repository.Db)).Find(users)
	return users, result.Error
}

func (repository *AuthRepositoryImpl) FindById(id string) (*model.User, error) {
	var user = &model.User{}
	result := repository.Db.Where("id = ?", id).Limit(1).Find(user)
	return user, result.Error
}

func (repository *AuthRepositoryImpl) FindByEmail(email string) (*model.User, error) {
	var user = &model.User{}
	result := repository.Db.Where(model.User{Email: email}).Where(model.User{Provider: ""}).Limit(1).Find(user)
	return user, result.Error
}

func (repository *AuthRepositoryImpl) FindByPhoneNumber(phoneNumber int) (*model.User, error) {
	var user = &model.User{}
	result := repository.Db.Where(model.User{PhoneNumber: phoneNumber}).Where(model.User{Provider: ""}).Limit(1).Find(user)
	return user, result.Error
}

func (repository *AuthRepositoryImpl) FindByProvider(provider string, providerUserId string) (*model.User, error) {
	var user = &model.User{}
	result := repository.Db.Where(model.User{Provider: provider}).Where(model.User{ProviderUserId: providerUserId}).Limit(1).Find(user)
	return user, result.Error
}
