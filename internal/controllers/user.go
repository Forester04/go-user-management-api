package controllers

import (
	"fmt"
	"net/http"

	"github.com/Forester04/go-user-management-api/internal/services"
	"github.com/Forester04/go-user-management-api/internal/viewmodel"
	"github.com/gin-gonic/gin"
)

func registerUserRoutes(group *gin.RouterGroup, svc services.ServiceInterface) {
	group.GET("/", listUsersController(svc))
	group.GET("/:id", requestViewmodelMiddleware(&viewmodel.GetUserRequest{}), getUserController(svc))
	group.DELETE("/:id", requestViewmodelMiddleware(&viewmodel.DeleteUserRequest{}), deleteUserController(svc))
}

func listUsersController(svc services.ServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := &viewmodel.GetAllUsersByEmailResponse{}

		users, err := svc.GetAllUsers()
		if err != nil {
			ctx.Error(err)
			return
		}

		response.Body.Emails = users

		ctx.Set(ContextKeyResponseviewmodel, response)

	}
}

func getUserController(svc services.ServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := ctx.MustGet(ContextKeyRequestViewmodel).(*viewmodel.GetUserRequest)
		response := &viewmodel.GetUserResponse{}

		user, err := svc.GetUser(request.ID)
		if err != nil {
			ctx.Error(err)
			return
		}

		response.Body.ID = user.ID
		response.Body.Name = fmt.Sprintf("%s %s", user.FirstName, user.LastName)

		ctx.Set(ContextKeyStatusCode, http.StatusOK)
		ctx.Set(ContextKeyResponseviewmodel, response)
	}
}

func deleteUserController(svc services.ServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := ctx.MustGet(ContextKeyRequestViewmodel).(*viewmodel.DeleteUserRequest)
		response := &viewmodel.DeleteUserResponse{}

		err := svc.DeleteUser(request.ID)
		if err != nil {
			ctx.Error(err)
			return
		}
		ctx.Set(ContextKeyStatusCode, http.StatusOK)
		ctx.Set(ContextKeyResponseviewmodel, response)
	}
}
