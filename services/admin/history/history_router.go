package history

import (
	"api/common/constants"
	"api/common/types"
	"api/services/admin"
	"api/services/admin/history/data"

	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *Controller,
) {
	var endpointConfig = types.ApiEndpointConfig{
		Group: "/history",
		Tag:   []string{"History"},
	}

	// Get all history
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-history-list",
			Summary:     "Get history list",
			Description: "Get history list with support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        endpointConfig.Group,
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						admin.FeaturePermission, // Feature scope
					},
				},
			},
			Metadata: map[string]any{
				constants.PermissionMetadata: constants.PermissionRead,
			},
			MaxBodyBytes:  1024, // 1 KiB
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
			Body data.HistoryList
		}, error) {
			result, errCode, err := controller.GetAll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.HistoryList
			}{Body: *result}, nil
		},
	)
}
