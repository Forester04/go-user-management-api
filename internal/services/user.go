package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Forester04/go-user-management-api/internal/dto"
	"github.com/Forester04/go-user-management-api/internal/errcode"
	"github.com/Forester04/go-user-management-api/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (svc *Service) RegisterUser(registerUser *dto.RegisterUser) (user *models.User, err error) {
	user, err = svc.globalRepository.User.GetByEmail(registerUser.Email)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrDatabase)
	}

	if user.ID != 0 {
		return nil, fmt.Errorf("%w", errcode.ErrUserAlreadyExist)
	}

	user, err = svc.formatRegisterUser(registerUser)
	if err != nil {
		return nil, err
	}

	err = svc.globalRepository.User.Create(user)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrDatabase, err)
	}
	return user, nil
}

func (svc *Service) LoginUser(email string, password string) (user *models.User, err error) {
	user, err = svc.globalRepository.User.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrDatabase, err)
	}

	if user == nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrInvalidCredentials, errors.New("user does not exist"))
	}

	// check if password is correct
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrInvalidCredentials, err)
	}
	return user, nil
}

func (svc *Service) formatRegisterUser(registerUser *dto.RegisterUser) (user *models.User, err error) {
	registerUser.Email = strings.ToLower(strings.TrimSpace(registerUser.Email))

	// password hashing
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(registerUser.Password), 12)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrExternalLib, err)
	}
	registerUser.Password = string(passwordHash)

	//parse birth date
	var parsedBirthDate time.Time
	if registerUser.BirthDate != "" {
		parsedBirthDate, err = time.Parse("2006-01-02", registerUser.BirthDate)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", errcode.ErrInvalidParameters, err)
		}
	}

	user = &models.User{
		Email:     registerUser.Email,
		Password:  registerUser.Password,
		FirstName: registerUser.FirstName,
		LastName:  registerUser.LastName,
		Phone:     &registerUser.Phone,
		BirthDate: &parsedBirthDate,
	}
	return user, nil
}
