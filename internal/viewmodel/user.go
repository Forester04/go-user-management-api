package viewmodel

type GetUserRequest struct {
	ID uint `json:"id"`
}

type GetUserResponse struct {
	Body struct {
		Id   uint   `json:"id"`
		Name string `json:"name"`
	} `json:"body"`
}

type DeleteUserRequest struct {
	ID uint `json:"id"`
}

type DeleteUserResponse struct {
	Body struct {
		ID uint `json:"id"`
	} `json:"body"`
}
