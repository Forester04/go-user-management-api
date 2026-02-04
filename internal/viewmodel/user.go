package viewmodel

import "github.com/Forester04/go-user-management-api/internal/models"

type GetUserRequest struct {
	ID uint `json:"id" uri:"id" binding:"required"`
}

type GetUserResponse struct {
	Body struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	} `json:"body"`
}

type GetAllUsersByEmailResponse struct {
	Body struct {
		Emails []*models.User `json:"emails"`
	} `json:"body"`
}

type DeleteUserRequest struct {
	ID uint `json:"id" uri:"id" binding:"required"`
}

type DeleteUserResponse struct {
	Body struct {
		ID uint `json:"id"`
	} `json:"body"`
}
