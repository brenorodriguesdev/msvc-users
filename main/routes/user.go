package routes

import (
	"msvc-users/main/adapters"
	"msvc-users/main/factories"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/signup", adapters.NewAdapterHttp(factories.MakeSignUpController()).Handle)
}
