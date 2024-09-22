package router

import (
	"github.com/4kpros/go-api/common/middleware"
	"github.com/gin-gonic/gin"
)

func GET(ginRouter *gin.RouterGroup, endpoint string, handler gin.HandlerFunc, requiredAuth bool, allowedRoles []string) {
	ginRouter.GET(endpoint, middleware.SecureAPIHandler(handler, requiredAuth, allowedRoles))
}

func POST(ginRouter *gin.RouterGroup, endpoint string, handler gin.HandlerFunc, requiredAuth bool, allowedRoles []string) {
	ginRouter.POST(endpoint, middleware.SecureAPIHandler(handler, requiredAuth, allowedRoles))
}

func PUT(ginRouter *gin.RouterGroup, endpoint string, handler gin.HandlerFunc, requiredAuth bool, allowedRoles []string) {
	ginRouter.PUT(endpoint, middleware.SecureAPIHandler(handler, requiredAuth, allowedRoles))
}

func PATCH(ginRouter *gin.RouterGroup, endpoint string, handler gin.HandlerFunc, requiredAuth bool, allowedRoles []string) {
	ginRouter.PATCH(endpoint, middleware.SecureAPIHandler(handler, requiredAuth, allowedRoles))
}

func DELETE(ginRouter *gin.RouterGroup, endpoint string, handler gin.HandlerFunc, requiredAuth bool, allowedRoles []string) {
	ginRouter.DELETE(endpoint, middleware.SecureAPIHandler(handler, requiredAuth, allowedRoles))
}
