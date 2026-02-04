package database

import (
	"fmt"

	"github.com/Forester04/go-user-management-api/internal/errcode"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func dnsBuilder() string {
	// host := viper.GetString("DATABASE_HOST")
	// port := viper.GetString("POSTGRES_PORT")
	// user := viper.GetString("POSTGRES_USER")
	// password := viper.GetString("POSTGRES_PASSWORD")
	// dbName := viper.GetString("POSTGRES_DB")
	connectionString := viper.GetString("CONNECTION_STRING")
	return connectionString
}

func NewGormClient() (*gorm.DB, error) {
	dns := dnsBuilder()

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrConfigurationFailed, err)
	}

	err = migrate(db)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errcode.ErrDatabaseMigration, err)
	}

	if viper.GetBool("SEED_DB") {
		//CreateEntities(db)
	}

	return db, nil
}
