package user

import (
	"fmt"
	"net/http"

	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/services/user/model"
)

type UserService interface {
	CreateWithEmail(user *model.User) (password string, errCode int, err error)
	CreateWithPhoneNumber(user *model.User) (password string, errCode int, err error)
	UpdateUser(user *model.User) (errCode int, err error)
	UpdateUserInfo(userInfo *model.UserInfo) (errCode int, err error)
	Delete(id string) (affectedRows int64, errCode int, err error)
	FindById(id string) (user *model.User, errCode int, err error)
	FindUserInfoById(id string) (user *model.UserInfo, errCode int, err error)
	FindAll(filter *types.Filter, pagination *types.Pagination) (users []model.User, errCode int, err error)
}

type UserServiceImpl struct {
	Repository UserRepository
}

func NewUserServiceImpl(repository UserRepository) UserService {
	return &UserServiceImpl{Repository: repository}
}

func (service *UserServiceImpl) CreateWithEmail(user *model.User) (password string, errCode int, err error) {
	// Check if user exists
	foundUser, errFound := service.Repository.FindByEmail(user.Email)
	if errFound != nil {
		errCode = http.StatusInternalServerError
		err = errFound
		return
	}
	if foundUser != nil && foundUser.Email == user.Email {
		message := "User with this email already exists! Please use another email."
		errCode = http.StatusFound
		err = fmt.Errorf("%s", message)
		return
	}

	// Create new user
	var randomPassword = utils.GenerateRandomPassword(8)
	foundUser.Email = user.Email
	foundUser.Password = randomPassword
	err = service.Repository.Create(user)
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}
	password = randomPassword
	return
}

func (service *UserServiceImpl) CreateWithPhoneNumber(user *model.User) (password string, errCode int, err error) {
	// Check if user exists
	foundUser, errFound := service.Repository.FindByPhoneNumber(user.PhoneNumber)
	if errFound != nil {
		errCode = http.StatusInternalServerError
		err = errFound
		return
	}
	if foundUser != nil && foundUser.PhoneNumber == user.PhoneNumber {
		message := "User with this phone number already exists! Please use another phone number."
		errCode = http.StatusFound
		err = fmt.Errorf("%s", message)
		return
	}

	// Create new user
	var randomPassword = utils.GenerateRandomPassword(8)
	foundUser.PhoneNumber = user.PhoneNumber
	foundUser.Password = randomPassword
	err = service.Repository.Create(user)
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}
	password = randomPassword
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
		message := "Could not delete user that doesn't exists! Please enter valid id."
		err = fmt.Errorf("%s", message)
		return
	}
	return
}

// FindById implements UserService.
func (service *UserServiceImpl) FindById(id string) (user *model.User, errCode int, err error) {
	user, err = service.Repository.FindById(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}
	if user == nil {
		errCode = http.StatusNotFound
		message := "User not found! Please enter valid id."
		err = fmt.Errorf("%s", message)
	}
	return
}

// FindUserInfoById implements UserService.
func (service *UserServiceImpl) FindUserInfoById(id string) (user *model.UserInfo, errCode int, err error) {
	user, err = service.Repository.FindUserInfoById(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}
	if user == nil {
		errCode = http.StatusNotFound
		message := "User information not found! Please enter valid id."
		err = fmt.Errorf("%s", message)
	}
	return
}

// FindAll implements UserService.
func (service *UserServiceImpl) FindAll(filter *types.Filter, pagination *types.Pagination) (users []model.User, errCode int, err error) {
	users, err = service.Repository.FindAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
	}
	return
}
