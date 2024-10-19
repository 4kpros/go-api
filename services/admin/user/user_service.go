package user

import (
	"api/common/constants"
	"api/common/types"
	"api/common/utils"
	"api/services/admin/user/model"
	"fmt"
	"net/http"
)

type UserService struct {
	Repository *UserRepository
}

func NewUserService(repository *UserRepository) *UserService {
	return &UserService{Repository: repository}
}

// Create user
func (service *UserService) Create(jwtToken *types.JwtToken, user *model.User) (result *model.User, errCode int, err error) {
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
		err = constants.HTTP_500_ERROR_MESSAGE(
			fmt.Sprintf("get user by %s from database", errMsg),
		)
		return
	}
	if foundUser != nil && foundUser.Email == user.Email {
		errCode = http.StatusFound
		err = constants.HTTP_302_ERROR_MESSAGE(
			fmt.Sprintf("user %s", errMsg),
		)
		return
	}

	// Create new user
	randomPassword := utils.GenerateRandomPassword(8)
	newUser := &model.User{
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		RoleId:       user.RoleId,
		Password:     randomPassword,
		SignInMethod: constants.AUTH_LOGIN_METHOD_DEFAULT,
	}
	result, err = service.Repository.Create(newUser)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("create user from database")
		return
	}
	return
}

// Update user
func (service *UserService) UpdateUser(jwtToken *types.JwtToken, user *model.User) (result *model.User, errCode int, err error) {
	result, err = service.Repository.UpdateUser(user.ID, user)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("update user from database")
	}
	return
}

// Delete user with matching id and return affected rows
func (service *UserService) Delete(jwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
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
func (service *UserService) Get(jwtToken *types.JwtToken, id int64) (user *model.User, errCode int, err error) {
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

// Return all users with support for search, filter and pagination
func (service *UserService) GetAll(jwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (userList []model.User, errCode int, err error) {
	userList, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("get users from database")
	}
	return
}
