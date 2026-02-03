package dto

type RegisterUser struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8,max=64"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Phone     string `json:"phone" binding:"omitempty,e164"`
	BirthDate string `json:"birth_date" binding:"omitempty,datetime=2006-01-02"`
	IsActive  bool   `json:"is_active"`
}
