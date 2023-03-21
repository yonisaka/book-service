package persistence

import (
	"context"

	"github.com/yonisaka/book-service/domain/entity"
	"github.com/yonisaka/book-service/domain/repository"

	"gorm.io/gorm"
)

type BookRepo struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepo {
	return &BookRepo{db}
}

var _ repository.BookRepositoryInterface = &BookRepo{}

func (u BookRepo) Create(ctx context.Context, book *entity.Book) error {
	return u.db.WithContext(ctx).Create(&book).Error
}

func (u BookRepo) Update(ctx context.Context, id int, book *entity.Book) error {
	return u.db.WithContext(ctx).Model(&entity.Book{}).Where("id = ?", id).Updates(&book).Error
}

func (u BookRepo) Find(ctx context.Context, id int) (*entity.Book, error) {
	var book entity.Book
	err := u.db.WithContext(ctx).First(&book, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (u BookRepo) Get(ctx context.Context) ([]entity.Book, error) {
	var books []entity.Book

	err := u.db.WithContext(ctx).Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (u BookRepo) Delete(ctx context.Context, id int) error {
	return u.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Book{}).Error
}