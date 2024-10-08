package middleware

import (
	"fmt"

	"github.com/danielgtaylor/huma/v2"
)

func ErrorsMiddleware(ctx huma.Context, next func(huma.Context)) {
	for err := range ctx.Operation().Errors {
		if err >= 400 {
			// huma.WriteErr(api, ctx, err, ctx.Context().Err().Error())
			fmt.Printf("API ERROR %s", ctx.Context().Err().Error())
			return
		}
	}

	next(ctx)
}
