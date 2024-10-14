package user

import (
	"github.com/4kpros/go-api/common/constants"
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

func (repository *UserRepository) Create(user *model.User) (*model.User, error) {
	result := *user
	return &result, repository.Db.Create(&result).Error
}

func (repository *UserRepository) UpdateUser(id int64, user *model.User) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			"email":        user.Email,
			"phone_number": user.PhoneNumber,
			"role_id":      user.RoleId,
		},
	).Error
}

func (repository *UserRepository) Delete(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.User{})
	return result.RowsAffected, result.Error
}

func (repository *UserRepository) GetById(id int64) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *UserRepository) GetByEmail(email string) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Where(
		"sign_in_method = ?", constants.AUTH_LOGIN_METHOD_DEFAULT,
	).Where(
		"email = ?", email,
	).Limit(1).Find(result).Error
}

func (repository *UserRepository) GetByPhoneNumber(phoneNumber uint64) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Where(
		"sign_in_method = ?", constants.AUTH_LOGIN_METHOD_DEFAULT,
	).Where(
		"phone_number = ?", phoneNumber,
	).Limit(1).Find(result).Error
}

func (repository *UserRepository) GetByProvider(provider string, providerUserId string) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Where(
		"sign_in_method = ?", constants.AUTH_LOGIN_METHOD_PROVIDER,
	).Where(
		"provider = ?", provider,
	).Where(
		"provider_user_id = ?", providerUserId,
	).Limit(1).Find(result).Error
}

func (repository *UserRepository) GetAll(filter *types.Filter, pagination *types.Pagination) ([]model.User, error) {
	result := []model.User{}
	return result, repository.Db.Scopes(utils.PaginationScope(result, pagination, filter, repository.Db)).Find(result).Error
}

// ----------------- Authentication service -----------------

func (repository *UserRepository) CreateUserInfo(userInfo *model.UserInfo) (*model.UserInfo, error) {
	result := *userInfo
	return &result, repository.Db.Create(&result).Error
}
func (repository *UserRepository) CreateUserMfa(userMfa *model.UserMfa) (*model.UserMfa, error) {
	result := *userMfa
	return &result, repository.Db.Create(&result).Error
}
func (repository *UserRepository) UpdateUserPassword(id int64, password string) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Model(result).Where("id = ?", id).Update("password", password).Error
}

func (repository *UserRepository) UpdateUserActivation(id int64, user *model.User) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			// "user_info_id": user.UserInfoId,
			// "user_mfa_id":  user.UserMfaId,
			"is_activated": user.IsActivated,
			"activated_at": user.ActivatedAt,
		},
	).Error
}

// ----------------- Profile service -----------------
func (repository *UserRepository) UpdateProfile(id int64, user *model.User) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			"email":        user.Email,
			"phone_number": user.PhoneNumber,
			"password":     user.Password,
		},
	).Error
}

func (repository *UserRepository) UpdateProfileInfo(id int64, userInfo *model.UserInfo) (*model.UserInfo, error) {
	result := &model.UserInfo{}
	return result, repository.Db.Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			"user_name":  userInfo.UserName,
			"first_name": userInfo.FirstName,
			"last_name":  userInfo.LastName,
			"address":    userInfo.Address,
			"image":      userInfo.Image,
			"language":   userInfo.Language,
		},
	).Error
}

func (repository *UserRepository) UpdateProfileMfa(id int64, column string, value bool) (*model.UserMfa, error) {
	result := &model.UserMfa{}
	return result, repository.Db.Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			"" + column: value,
		},
	).Error
}
