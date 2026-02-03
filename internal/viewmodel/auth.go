package viewmodel

import "github.com/Forester04/go-user-management-api/internal/dto"

type RegisteruserRequest struct {
	Body dto.RegisterUser `json:"body" binding:"required"`
}

type RegisterUserResponse struct {
	Body struct {
		Token string `json:"token"`
	} `json:"body"`
}

type LoginUserRequest struct {
	Body struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8,max=64"`
	} `json:"body" binding:"required"`
}

type LoginUserResponse struct {
	Body struct {
		Token string `json:"token"`
	} `json:"body"`
}
