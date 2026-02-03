package repositories

import (
	"gorm.io/gorm"
)

type GlobalRepository struct {
	User UserRepositoryInterface
}

func NewGlobalRepository(DB *gorm.DB) *GlobalRepository {
	gr := &GlobalRepository{
		User: &UserRepository{DB: DB},
	}
	return gr
}
