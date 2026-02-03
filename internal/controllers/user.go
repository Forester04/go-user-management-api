package controllers

import (
	"fmt"
	"net/http"

	"github.com/Forester04/go-user-management-api/internal/services"
	"github.com/Forester04/go-user-management-api/internal/viewmodel"
	"github.com/gin-gonic/gin"
)

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

func deleteUser(svc services.ServiceInterface) gin.HandlerFunc {
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
