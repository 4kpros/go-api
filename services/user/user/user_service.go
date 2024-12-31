package user

import (
	"fmt"
	"net/http"
	"time"

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
func (service *Service) Create(inputJwtToken *types.JwtToken, user *model.User) (result *model.User, errCode int, err error) {
	// Check if user exists
	var foundUser *model.User
	var errMsg string = ""
	var isEmailValid = utils.IsEmailValid(user.Email)
	if isEmailValid {
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
	if foundUser != nil {
		if (isEmailValid && foundUser.Email == user.Email) || (!isEmailValid && foundUser.PhoneNumber == user.PhoneNumber) {
			errCode = http.StatusFound
			err = constants.Http302ErrorMessage(
				fmt.Sprintf("user %s", errMsg),
			)
			return
		}
	}

	// Create new user
	randomPassword := utils.GenerateRandomPassword(8)
	var activatedAt *time.Time = nil
	if user.IsActivated {
		tmpTime := time.Now()
		activatedAt = &tmpTime
	}
	newUser := &model.User{
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		RoleID:      user.RoleID,
		IsActivated: user.IsActivated,
		ActivatedAt: activatedAt,
		LoginMethod: constants.AuthLoginMethodDefault,
		Password:    randomPassword,
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
func (service *Service) Update(inputJwtToken *types.JwtToken, userID int64, user *model.User) (result *model.User, errCode int, err error) {
	// Check if user exists
	var errMsg string = ""
	errMsg = "email"
	foundUser, err := service.Repository.GetByEmail(user.Email)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(
			fmt.Sprintf("get user by %s from database", errMsg),
		)
		return
	}
	if foundUser != nil {
		if foundUser.Email != user.Email {
			errCode = http.StatusFound
			err = constants.Http302ErrorMessage(
				fmt.Sprintf("user %s", errMsg),
			)
			return
		}
	}

	errMsg = "phone number"
	foundUser, err = service.Repository.GetByPhoneNumber(user.PhoneNumber)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(
			fmt.Sprintf("get user by %s from database", errMsg),
		)
		return
	}
	if foundUser != nil {
		if foundUser.PhoneNumber != user.PhoneNumber {
			errCode = http.StatusFound
			err = constants.Http302ErrorMessage(
				fmt.Sprintf("user %s", errMsg),
			)
			return
		}
	}

	// Update
	if !user.IsActivated {
		user.ActivatedAt = nil
	} else if user.IsActivated && !foundUser.IsActivated {
		tmpTime := time.Now()
		user.ActivatedAt = &tmpTime
	}
	result, err = service.Repository.Update(userID, user)
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

// Delete Deletes selection
func (service *Service) DeleteMultiple(inputJwtToken *types.JwtToken, list []int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultiple(list)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete multiple user from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("User selection")
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
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, roleName string) (userList []model.User, errCode int, err error) {
	userList, err = service.Repository.GetAll(filter, pagination, roleName)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get users from database")
	}
	return
}
