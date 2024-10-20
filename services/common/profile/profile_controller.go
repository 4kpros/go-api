package profile

import (
	"context"

	"api/common/helpers"
	"api/services/admin/user/model"
	"api/services/common/profile/data"
)

type Controller struct {
	Service *Service
}

func NewProfileController(service *Service) *Controller {
	return &Controller{Service: service}
}

func (controller *Controller) UpdateProfileEmail(
	ctx *context.Context,
	input *struct {
		Body data.UpdateProfileRequest
	},
) (result *model.User, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateProfileEmail(
		helpers.GetJwtContext(ctx),
		&model.User{
			Email:       input.Body.Email,
			PhoneNumber: input.Body.PhoneNumber,
			Password:    input.Body.Password,
		},
	)
	return
}

func (controller *Controller) UpdateProfileInfo(
	ctx *context.Context,
	input *struct {
		Body data.UpdateProfileInfoRequest
	},
) (result *model.UserInfo, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateProfileInfo(
		helpers.GetJwtContext(ctx),
		&model.UserInfo{
			UserName:  input.Body.UserName,
			FirstName: input.Body.FirstName,
			LastName:  input.Body.LastName,
			Address:   input.Body.Address,
			Image:     input.Body.Image,
			Language:  input.Body.Language,
		},
	)
	return
}

func (controller *Controller) UpdateProfileMfa(
	ctx *context.Context,
	input *struct {
		Body data.UpdateProfileMfaRequest
	},
) (result *model.UserMfa, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateProfileMfa(helpers.GetJwtContext(ctx), input.Body.Method, input.Body.Value)
	return
}

func (controller *Controller) DeleteProfile(
	ctx *context.Context,
	input *struct{},
) (result int64, errCode int, err error) {
	result, errCode, err = controller.Service.DeleteProfile(helpers.GetJwtContext(ctx))
	return
}

func (controller *Controller) GetProfile(
	ctx *context.Context,
	input *struct{},
) (result *model.User, errCode int, err error) {
	user, errCode, err := controller.Service.GetProfile(helpers.GetJwtContext(ctx))
	if err != nil {
		return
	}
	result = user
	return
}
