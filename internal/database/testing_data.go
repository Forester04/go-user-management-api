package database

import "github.com/Forester04/go-user-management-api/internal/models"

var (
	DummyUsers = []models.User{
		{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "johndoe@gmail.com",
			Password:  "john123",
		},
		{
			FirstName: "Paul",
			LastName:  "Dohn",
			Email:     "pauldohn@gmail.com",
			Password:  "paul123",
		},
	}
)
