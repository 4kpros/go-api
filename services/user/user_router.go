package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/4kpros/go-api/common/middleware"
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/services/user/data"
	"github.com/4kpros/go-api/services/user/model"
	"github.com/danielgtaylor/huma/v2"
)

func SetupEndpoints(
	humaApi *huma.API,
	controller *UserController,
) {
	var endpointConfig = struct {
		Group string
		Tag   []string
	}{
		Group: "/users",
		Tag:   []string{"Users"},
	}
	const requireAuth = true

	// Create user with email
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "post-user-email",
			Summary:       "Create user with email",
			Description:   "Create new user by providing email and user and return created object.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/email", endpointConfig.Group),
			Middlewares:   *middleware.GenerateMiddlewares(requireAuth),
			Tags:          endpointConfig.Tag,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.UserWithEmailRequest
			},
		) (*struct{ Body model.User }, error) {
			result, errCode, err := controller.CreateWithEmail(&input.Body)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body model.User }{Body: *result}, nil
		},
	)

	// Create user with phone number
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "create-user-phone",
			Summary:       "Create user with phone",
			Description:   "Create new user by providing phone number and role and return created object.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/phone", endpointConfig.Group),
			Middlewares:   *middleware.GenerateMiddlewares(requireAuth),
			Tags:          endpointConfig.Tag,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.UserWithPhoneNumberRequest
			},
		) (*struct{ Body model.User }, error) {
			result, errCode, err := controller.CreateWithPhoneNumber(&input.Body)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body model.User }{Body: *result}, nil
		},
	)

	// Update user with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "update-user",
			Summary:       "Update user",
			Description:   "Update existing user with matching id and return the new user object.",
			Method:        http.MethodPut,
			Path:          fmt.Sprintf("%s/:id", endpointConfig.Group),
			Middlewares:   *middleware.GenerateMiddlewares(requireAuth),
			Tags:          endpointConfig.Tag,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				data.UserId
				Body data.UserRequest
			},
		) (*struct{ Body model.User }, error) {
			var inputFormatted = &model.User{
				Email:       input.Body.Email,
				PhoneNumber: input.Body.PhoneNumber,
				Language:    input.Body.Language,
				Role:        input.Body.Role,
			}
			inputFormatted.ID = uint(input.UserId.Id)
			result, errCode, err := controller.UpdateUser(inputFormatted)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body model.User }{Body: *result}, nil
		},
	)

	// Update user info with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "update-user-info",
			Summary:       "Update user info",
			Description:   "Update existing user info with matching id and return the new user object.",
			Method:        http.MethodPut,
			Path:          fmt.Sprintf("%s/info/:id", endpointConfig.Group),
			Middlewares:   *middleware.GenerateMiddlewares(requireAuth),
			Tags:          endpointConfig.Tag,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				data.UserId
				Body data.UserInfoRequest
			},
		) (*struct{ Body model.UserInfo }, error) {
			var inputFormatted = &model.UserInfo{
				UserName:  input.Body.UserName,
				FirstName: input.Body.FirstName,
				LastName:  input.Body.LastName,
				Address:   input.Body.Address,
			}
			inputFormatted.ID = uint(input.UserId.Id)
			result, errCode, err := controller.UpdateUserInfo(inputFormatted)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body model.UserInfo }{Body: *result}, nil
		},
	)

	// Delete user with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "delete-user",
			Summary:       "Delete user",
			Description:   "Delete existing user with matching id and return affected rows in database.",
			Method:        http.MethodDelete,
			Path:          fmt.Sprintf("%s/:id", endpointConfig.Group),
			Middlewares:   *middleware.GenerateMiddlewares(requireAuth),
			Tags:          endpointConfig.Tag,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				data.UserId
			},
		) (*struct{ Body types.DeleteResponse }, error) {
			result, errCode, err := controller.Delete(&input.UserId)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeleteResponse }{Body: types.DeleteResponse{AffectedRows: result}}, nil
		},
	)

	// Find user by id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "find-user-id",
			Summary:       "Find user by id",
			Description:   "Return one user with matching id",
			Method:        http.MethodGet,
			Path:          fmt.Sprintf("%s/:id", endpointConfig.Group),
			Middlewares:   *middleware.GenerateMiddlewares(requireAuth),
			Tags:          endpointConfig.Tag,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				data.UserId
			},
		) (*struct{ Body model.User }, error) {
			result, errCode, err := controller.FindById(&input.UserId)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body model.User }{Body: *result}, nil
		},
	)

	// Find all users
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "find-users",
			Summary:       "Find all users",
			Description:   "Find all users with support for search, filter and pagination",
			Method:        http.MethodGet,
			Path:          fmt.Sprintf("%s ", endpointConfig.Group),
			Middlewares:   *middleware.GenerateMiddlewares(requireAuth),
			Tags:          endpointConfig.Tag,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden},
		},
		func(
			ctx context.Context,
			input *struct {
				types.Filter
				types.PaginationRequest
			},
		) (*struct {
			Body data.UsersResponse
		}, error) {
			result, errCode, err := controller.FindAll(&input.Filter, &input.PaginationRequest)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.UsersResponse
			}{Body: *result}, nil
		},
	)
}
