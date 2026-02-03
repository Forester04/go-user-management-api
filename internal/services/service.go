package services

import (
	"github.com/Forester04/go-user-management-api/internal/repositories"
	"go.uber.org/zap"
)

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
