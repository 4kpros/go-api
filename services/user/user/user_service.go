package user

import (
	"fmt"
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/common/utils"
	"api/services/user/user/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Create user
func (service *Service) Create(inputJwtToken *types.JwtToken, roleID int64, user *model.User) (result *model.User, errCode int, err error) {
	// Check if user exists
	var foundUser *model.User
	var errMsg string = ""
	if utils.IsEmailValid(user.Email) {
		errMsg = "email"
		foundUser, err = service.Repository.GetByEmail(user.Email)
	} else {
		errMsg = "phone number"
		foundUser, err = service.Repository.GetByPhoneNumber(user.PhoneNumber)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(
			fmt.Sprintf("get user by %s from database", errMsg),
		)
		return
	}
	if foundUser != nil && foundUser.Email == user.Email {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(
			fmt.Sprintf("user %s", errMsg),
		)
		return
	}

	// Create new user
	randomPassword := utils.GenerateRandomPassword(8)
	newUser := &model.User{
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Password:    randomPassword,
		LoginMethod: constants.AuthLoginMethodDefault,
	}
	result, err = service.Repository.Create(newUser)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create user from database")
		return
	}
	return
}

// CreateUserRole assign role to user
func (service *Service) AssignRole(inputJwtToken *types.JwtToken, userID int64, roleID int64) (result *model.User, errCode int, err error) {
	result, err = service.Repository.AssignRole(userID, roleID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("assign role to user from database")
		return
	}
	return
}

// UpdateUser Update user
func (service *Service) Update(inputJwtToken *types.JwtToken, roleID int64, user *model.User) (result *model.User, errCode int, err error) {
	// Check if user exists
	var foundUser *model.User
	var errMsg string = ""
	if utils.IsEmailValid(user.Email) {
		errMsg = "email"
		foundUser, err = service.Repository.GetByEmail(user.Email)
	} else {
		errMsg = "phone number"
		foundUser, err = service.Repository.GetByPhoneNumber(user.PhoneNumber)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(
			fmt.Sprintf("get user by %s from database", errMsg),
		)
		return
	}
	if foundUser != nil && foundUser.Email == user.Email {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(
			fmt.Sprintf("user %s", errMsg),
		)
		return
	}

	// Update
	result, err = service.Repository.Update(user.ID, user)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update user from database")
	}
	return
}

// Delete user with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, userID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(userID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete user from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("User")
		return
	}
	return
}

// DeleteUserRole remove user role and return affected rows
func (service *Service) DeleteRole(inputJwtToken *types.JwtToken, userID int64, roleID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteRole(userID, roleID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete user role from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("User role")
		return
	}
	return
}

// Get Returns user with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, userID int64) (user *model.User, errCode int, err error) {
	user, err = service.Repository.GetByID(userID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get user by id from database")
		return
	}
	if user == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("User")
		return
	}
	return
}

// GetAll Returns all users with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (userList []model.User, errCode int, err error) {
	userList, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get users from database")
	}
	return
}
