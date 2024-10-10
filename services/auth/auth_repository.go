package auth

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/services/user/model"
	"gorm.io/gorm"
)

type AuthRepository struct {
	Db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{Db: db}
}

func (repository *AuthRepository) Create(user *model.User) error {
	var result = repository.Db.Create(user)
	return result.Error
}

func (repository *AuthRepository) CreateUserInfo(userInfo *model.UserInfo) error {
	var result = repository.Db.Create(userInfo)
	return result.Error
}

func (repository *AuthRepository) Update(user *model.User) error {
	var result = repository.Db.Model(user).Updates(user)
	return result.Error
}

func (repository *AuthRepository) UpdatePasswordById(id int64, password string) (*model.User, error) {
	var user = &model.User{
		Password: password,
	}
	var result = repository.Db.Model(user).Where("id = ?", id).Update("password", user)
	return user, result.Error
}

func (repository *AuthRepository) Delete(id int64) (int64, error) {
	var user = &model.User{}
	var result = repository.Db.Where("id = ?", id).Delete(user)
	return result.RowsAffected, result.Error
}

func (repository *AuthRepository) GetAll(filter *types.Filter, pagination *types.Pagination) ([]model.User, error) {
	var users = []model.User{}
	var result = repository.Db.Scopes(utils.PaginationScope(users, pagination, filter, repository.Db)).Find(users)
	return users, result.Error
}

func (repository *AuthRepository) GetById(id int64) (*model.User, error) {
	var user = &model.User{}
	var result = repository.Db.Where("id = ?", id).Limit(1).Find(user)
	return user, result.Error
}

func (repository *AuthRepository) GetByEmail(email string) (*model.User, error) {
	var user = &model.User{}
	var result = repository.Db.Where(model.User{Email: email}).Where(model.User{Provider: ""}).Limit(1).Find(user)
	return user, result.Error
}

func (repository *AuthRepository) GetByPhoneNumber(phoneNumber uint64) (*model.User, error) {
	var user = &model.User{}
	var result = repository.Db.Where(model.User{PhoneNumber: phoneNumber}).Where(model.User{Provider: ""}).Limit(1).Find(user)
	return user, result.Error
}

func (repository *AuthRepository) GetByProvider(provider string, providerUserId string) (*model.User, error) {
	var user = &model.User{}
	var result = repository.Db.Where(model.User{Provider: provider}).Where(model.User{ProviderUserId: providerUserId}).Limit(1).Find(user)
	return user, result.Error
}
