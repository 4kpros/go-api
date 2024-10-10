package user

import (
	"fmt"
	"net/http"

	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/services/user/model"
)

type UserService interface {
	Create(user *model.User) (result *model.User, errCode int, err error)
	UpdateUser(user *model.User) (errCode int, err error)
	UpdateUserInfo(userInfo *model.UserInfo) (errCode int, err error)
	Delete(id string) (result int64, errCode int, err error)
	GetById(id string) (result *model.User, errCode int, err error)
	GetUserInfoById(id string) (result *model.UserInfo, errCode int, err error)
	GetAll(filter *types.Filter, pagination *types.Pagination) (result []model.User, errCode int, err error)
}

type UserServiceImpl struct {
	Repository UserRepository
}

func NewUserServiceImpl(repository UserRepository) UserService {
	return &UserServiceImpl{Repository: repository}
}

func (service *UserServiceImpl) Create(user *model.User) (result *model.User, errCode int, err error) {
	// Check if user exists
	var foundUser *model.User = nil
	var errFound error = nil
	var message string = ""
	if utils.IsEmailValid(user.Email) {
		message = "User with this email already exists! Please use another email."
		foundUser, errFound = service.Repository.GetByEmail(user.Email)
	} else {
		message = "User with this phone number already exists! Please use another phone number."
		foundUser, errFound = service.Repository.GetByPhoneNumber(user.PhoneNumber)
	}
	if errFound != nil {
		errCode = http.StatusInternalServerError
		err = errFound
		return
	}
	if foundUser != nil && foundUser.Email == user.Email {
		errCode = http.StatusFound
		err = fmt.Errorf("%s", message)
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
		return
	}
	result = newUser
	return
}

func (service *UserServiceImpl) UpdateUser(user *model.User) (errCode int, err error) {
	err = service.Repository.UpdateUser(user)
	if err != nil {
		errCode = http.StatusInternalServerError
	}
	return
}

func (service *UserServiceImpl) UpdateUserInfo(userInfo *model.UserInfo) (errCode int, err error) {
	err = service.Repository.UpdateUserInfo(userInfo)
	if err != nil {
		errCode = http.StatusInternalServerError
	}
	return
}

func (service *UserServiceImpl) Delete(id string) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(id)
	if err != nil {
		errCode = http.StatusInternalServerError
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		var message = "Could not delete user that doesn't exists! Please enter valid id."
		err = fmt.Errorf("%s", message)
		return
	}
	return
}

func (service *UserServiceImpl) GetById(id string) (user *model.User, errCode int, err error) {
	user, err = service.Repository.GetById(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}
	if user == nil {
		errCode = http.StatusNotFound
		var message = "User not found! Please enter valid id."
		err = fmt.Errorf("%s", message)
	}
	return
}

func (service *UserServiceImpl) GetUserInfoById(id string) (user *model.UserInfo, errCode int, err error) {
	user, err = service.Repository.GetUserInfoById(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}
	if user == nil {
		errCode = http.StatusNotFound
		var message = "User information not found! Please enter valid id."
		err = fmt.Errorf("%s", message)
	}
	return
}

func (service *UserServiceImpl) GetAll(filter *types.Filter, pagination *types.Pagination) (users []model.User, errCode int, err error) {
	users, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
	}
	return
}
