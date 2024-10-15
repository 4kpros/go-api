package profile

import (
	"context"

	"github.com/4kpros/go-api/common/helpers"
	"github.com/4kpros/go-api/services/user/data"
	"github.com/4kpros/go-api/services/user/model"
)

type ProfileController struct {
	Service *ProfileService
}

func NewProfileController(service *ProfileService) *ProfileController {
	return &ProfileController{Service: service}
}

func (controller *ProfileController) UpdateProfile(
	ctx *context.Context,
	input *struct {
		Body data.UpdateProfileRequest
	},
) (result *model.User, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateProfile(
		helpers.GetJwtContext(ctx),
		&model.User{
			Email:       input.Body.Email,
			PhoneNumber: input.Body.PhoneNumber,
			Password:    input.Body.Password,
		},
	)
	return
}

func (controller *ProfileController) UpdateProfileInfo(
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

func (controller *ProfileController) UpdateProfileMfa(
	ctx *context.Context,
	input *struct {
		Body data.UpdateProfileMfaRequest
	},
) (result *model.UserMfa, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateProfileMfa(helpers.GetJwtContext(ctx), input.Body.Method, input.Body.Value)
	return
}

func (controller *ProfileController) DeleteProfile(
	ctx *context.Context,
	input *struct{},
) (result int64, errCode int, err error) {
	result, errCode, err = controller.Service.DeleteProfile(helpers.GetJwtContext(ctx))
	return
}

func (controller *ProfileController) GetProfile(
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
