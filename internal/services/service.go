package services

import (
	"github.com/Forester04/go-user-management-api/internal/dto"
	"github.com/Forester04/go-user-management-api/internal/models"
	"github.com/Forester04/go-user-management-api/internal/repositories"
	"go.uber.org/zap"
)

type ServiceInterface interface {
	RegisterUser(registerUser *dto.RegisterUser) (user *models.User, err error)
	LoginUser(email string, password string) (user *models.User, err error)
	formatRegisterUser(registerUser *dto.RegisterUser) (user *models.User, err error)
	GenerateToken(user *models.User) (toeknString string, err error)
	GetUser(id uint) (user *models.User, err error)
	DeleteUser(id uint) (err error)
}

type Service struct {
	logger           *zap.Logger
	globalRepository *repositories.GlobalRepository
}

func New(logger *zap.Logger, globalRepository *repositories.GlobalRepository) *Service {
	service := &Service{
		logger,
		globalRepository,
	}
	return service
}
