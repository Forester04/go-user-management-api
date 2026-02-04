package main

import (
	"fmt"
	"os"

	"github.com/Forester04/go-user-management-api/internal/controllers"
	"github.com/Forester04/go-user-management-api/internal/database"
	"github.com/Forester04/go-user-management-api/internal/logger"
	"github.com/Forester04/go-user-management-api/internal/repositories"
	"github.com/Forester04/go-user-management-api/internal/services"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	initViper()

	//Initialize logger
	logger, err := logger.New()
	if err != nil {
		panic(fmt.Errorf("error logger: %w", err))
	}
	defer func(Logger *zap.Logger) {
		// Sync logger before exit
		if err := Logger.Sync(); err != nil {
			logger.Fatal("error syncing logger", zap.Error(err))
		}
	}(logger)

	//Initialize database
	gormClient, err := database.NewGormClient()
	if err != nil {
		logger.Fatal("error initializing database", zap.Error(err))
	}

	// Initializing respositories
	globalRepository := repositories.NewGlobalRepository(gormClient)

	// Initializing services
	service := services.New(logger, globalRepository)

	//Initialize handlers
	routing := controllers.NewRouter(logger, service)
	err = routing.Run(":" + viper.GetString("PORT"))
	if err != nil {
		logger.Fatal("error running router", zap.Error(err))
	}

}

func initViper() {
	//Load environment variables
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error config file: %w", err))
	}
	if viper.GetString("HTTP_PROXY") != "" {
		os.Setenv("HTTP_PROXY", viper.GetString("HTTP_PROXY"))
	}
	if viper.GetString("HTTPS_PROXY") != "" {
		os.Setenv("HTTPS_PROXY", viper.GetString("HTTP_PROXY"))
	}
}
