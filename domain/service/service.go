package service

import (
	"github.com/yonisaka/book-service/domain/repository"
	"github.com/yonisaka/book-service/infrastructure/persistence"
	"gorm.io/gorm"
)

// Repositories is a struct
type Repositories struct {
	Book    repository.BookRepositoryInterface
	HttpLog repository.HttpLogRepositoryInterface
	DB      *gorm.DB
}

// NewDBService is constructor
func NewDBService(db *gorm.DB) *Repositories {
	return &Repositories{
		Book:    persistence.NewBookRepository(db),
		HttpLog: persistence.NewHttpLogRepository(db),
		DB:      db,
	}
}
