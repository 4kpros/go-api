package history

import (
	"context"
	"fmt"
	"net/http"

	"github.com/4kpros/go-api/common/middleware"
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/services/history/data"
	"github.com/danielgtaylor/huma/v2"
)

func SetupEndpoints(
	humaApi *huma.API,
	controller *HistoryController,
) {
	var endpointConfig = struct {
		Group string
		Tag   []string
	}{
		Group: "/history",
		Tag:   []string{"History"},
	}
	const requireAuth = true

	// Get all history
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "get-history",
			Summary:       "Get history",
			Description:   "Get history with support for search, filter and pagination",
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
			Body data.HistoriesResponse
		}, error) {
			var result, errCode, err = controller.GetAll(&input.Filter, &input.PaginationRequest)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.HistoriesResponse
			}{Body: *result}, nil
		},
	)
}
