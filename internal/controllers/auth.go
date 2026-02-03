package controllers

import (
	"net/http"

	"github.com/Forester04/go-user-management-api/internal/services"
	"github.com/Forester04/go-user-management-api/internal/viewmodel"
	"github.com/gin-gonic/gin"
)

func registerAuthRoutes(group *gin.RouterGroup, svc services.ServiceInterface) {
	group.POST("/register", requestViewmodelMiddleware(&viewmodel.RegisteruserRequest{}), registerController(svc))
	group.POST("/login", requestViewmodelMiddleware(&viewmodel.LoginUserRequest{}), loginController(svc))
}

func registerController(svc services.ServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := ctx.MustGet(ContextKeyRequestViewmodel).(*viewmodel.RegisteruserRequest)
		response := &viewmodel.RegisterUserResponse{}

		user, err := svc.RegisterUser(&request.Body)
		if err != nil {
			ctx.Error(err)
			return
		}

		token, err := svc.GenerateToken(user)
		if err != nil {
			ctx.Error(err)
			return
		}

		response.Body.Token = token

		ctx.Set(ContextKeyStatusCode, http.StatusOK)
		ctx.Set(ContextKeyResponseviewmodel, response)
	}
}

func loginController(svc services.ServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := ctx.MustGet(ContextKeyRequestViewmodel).(*viewmodel.LoginUserRequest)
		response := &viewmodel.LoginUserResponse{}

		user, err := svc.LoginUser(request.Body.Email, request.Body.Password)
		if err != nil {
			ctx.Error(err)
			return
		}

		token, err := svc.GenerateToken(user)
		if err != nil {
			ctx.Error(err)
			return
		}

		response.Body.Token = token

		ctx.Set(ContextKeyStatusCode, http.StatusOK)
		ctx.Set(ContextKeyResponseviewmodel, response)
	}
}
