package user

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/services/user/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (repository *UserRepository) Create(user *model.User) error {
	return repository.Db.Create(user).Error
}

func (repository *UserRepository) UpdateUser(user *model.User) error {
	return repository.Db.Model(user).Updates(user).Error
}

func (repository *UserRepository) UpdateUserInfo(userInfo *model.UserInfo) error {
	return repository.Db.Model(userInfo).Updates(userInfo).Error
}

func (repository *UserRepository) Delete(id int64) (int64, error) {
	user := &model.User{}
	result := repository.Db.Where("id = ?", id).Delete(user)
	return result.RowsAffected, result.Error
}

func (repository *UserRepository) GetById(id int64) (*model.User, error) {
	user := &model.User{}
	result := repository.Db.Where("id = ?", id).Limit(1).Find(user)
	return user, result.Error
}

func (repository *UserRepository) GetUserInfoById(id int64) (*model.UserInfo, error) {
	userInfo := &model.UserInfo{}
	result := repository.Db.Where("id = ?", id).Limit(1).Find(userInfo)
	return userInfo, result.Error
}

func (repository *UserRepository) GetByEmail(email string) (*model.User, error) {
	user := &model.User{}
	result := repository.Db.Where("email = ? AND (provider is null OR provider = '')", email).Limit(1).Find(user)
	return user, result.Error
}

func (repository *UserRepository) GetByPhoneNumber(phoneNumber uint64) (*model.User, error) {
	user := &model.User{}
	result := repository.Db.Where("phoneNumber = ? AND (provider is null OR provider = '')", phoneNumber).Limit(1).Find(user)
	return user, result.Error
}

func (repository *UserRepository) GetByProvider(provider string, providerUserId string) (*model.User, error) {
	user := &model.User{}
	result := repository.Db.Where("provider = ? AND providerUserId = ?", provider, providerUserId).Limit(1).Find(user)
	return user, result.Error
}

func (repository *UserRepository) GetAll(filter *types.Filter, pagination *types.Pagination) ([]model.User, error) {
	userList := []model.User{}
	result := repository.Db.Scopes(utils.PaginationScope(userList, pagination, filter, repository.Db)).Find(userList)
	return userList, result.Error
}
