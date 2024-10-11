package user

import (
	"net/http"

	"github.com/4kpros/go-api/common/constants"
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/services/user/model"
)

type UserService struct {
	Repository *UserRepository
}

func NewUserService(repository *UserRepository) *UserService {
	return &UserService{Repository: repository}
}

// Create user
func (service *UserService) Create(user *model.User) (result *model.User, errCode int, err error) {
	// Check if user exists
	var foundUser *model.User = nil
	var errFound error = nil
	var message string = ""
	if utils.IsEmailValid(user.Email) {
		message = "user email"
		foundUser, errFound = service.Repository.GetByEmail(user.Email)
	} else {
		message = "user phone number"
		foundUser, errFound = service.Repository.GetByPhoneNumber(user.PhoneNumber)
	}
	if errFound != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("get user by email/phone from database")
		return
	}
	if foundUser != nil && foundUser.Email == user.Email {
		errCode = http.StatusFound
		err = constants.HTTP_302_ERROR_MESSAGE(message)
		return
	}

	// Create new user
	var randomPassword = utils.GenerateRandomPassword(8)
	var newUser = &model.User{
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Password:    randomPassword,
		RoleId:      user.RoleId,
	}
	err = service.Repository.Create(newUser)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("create user from database")
		return
	}
	result = newUser
	return
}

// Update user
func (service *UserService) UpdateUser(user *model.User) (errCode int, err error) {
	err = service.Repository.UpdateUser(user)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("update user from database")
	}
	return
}

// Update info
func (service *UserService) UpdateUserInfo(userInfo *model.UserInfo) (errCode int, err error) {
	err = service.Repository.UpdateUserInfo(userInfo)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("update user info from database")
	}
	return
}

// Delete user with matching id and return affected rows
func (service *UserService) Delete(id int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("delete user from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.HTTP_404_ERROR_MESSAGE("User")
		return
	}
	return
}

// Return user with matching id
func (service *UserService) GetById(id int64) (user *model.User, errCode int, err error) {
	user, err = service.Repository.GetById(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("get user by id from database")
		return
	}
	if user == nil {
		errCode = http.StatusNotFound
		err = constants.HTTP_404_ERROR_MESSAGE("User")
		return
	}
	return
}

// Return user info with matching id
func (service *UserService) GetUserInfoById(id int64) (userInfo *model.UserInfo, errCode int, err error) {
	userInfo, err = service.Repository.GetUserInfoById(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("get user info by id from database")
		return
	}
	if userInfo == nil {
		errCode = http.StatusNotFound
		err = constants.HTTP_404_ERROR_MESSAGE("User info")
		return
	}
	return
}

// Return all users with support for search, filter and pagination
func (service *UserService) GetAll(filter *types.Filter, pagination *types.Pagination) (users []model.User, errCode int, err error) {
	users, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("get users from database")
	}
	return
}
