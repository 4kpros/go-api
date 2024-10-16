package admin

import (
	"context"
	"net/http"

	"api/common/types"
	"api/services/admin/data"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *AdminController,
) {
	var endpointConfig = types.APIEndpointConfig{
		Group: "/admin",
		Tag:   []string{"Admin"},
	}

	// Create admin
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "post-admin",
			Summary:       "Create admin",
			Description:   "Create new admin by providing the token, email and password. After that this endpoint will no longer work",
			Method:        http.MethodPost,
			Path:          endpointConfig.Group,
			Tags:          endpointConfig.Tag,
			MaxBodyBytes:  1024, // 1 KiB
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.CreateAdminRequest
			},
		) (*struct{ Body data.CreateAdminResponse }, error) {
			_, errCode, err := controller.Create(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.CreateAdminResponse }{
				Body: data.CreateAdminResponse{
					Message: "Admin created! Let's now connect.",
				},
			}, nil
		},
	)
}
