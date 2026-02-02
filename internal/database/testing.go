package database

import (
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetFailedTx(db *gorm.DB) *gorm.DB {
	failedTx := db.Begin()
	failedTx.Rollback()
	return failedTx
}

func CreateEntities(db *gorm.DB) {
	dummyEntities := []interface{}{
		DummyUsers,
	}

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := createEntities(tx, dummyEntities...); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatalf("Failed to create entities %v", err)
	}
}

func createEntities(tx *gorm.DB, entities ...interface{}) error {
	for _, entity := range entities {
		// check if table is empty
		res := tx.Model(entity).First(&struct{}{})
		if res.Error == nil || res.Error != gorm.ErrRecordNotFound {
			fmt.Println("Table not empty, skipping...")
		}

		// create entity if the table is empty
		if err := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}},
		}).Create(entity).Error; err != nil {
			if gorm.ErrDuplicatedKey == err {
				fmt.Printf("%T already exist\n", entity)
				continue
			}
			return err
		}
	}
	return nil
}
