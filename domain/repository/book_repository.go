package repository

import (
	"context"

	"github.com/yonisaka/book-service/domain/entity"
)

type BookRepositoryInterface interface {
	Create(ctx context.Context, book *entity.Book) error
	Update(ctx context.Context, id int, book *entity.Book) error
	Find(ctx context.Context, id int) (*entity.Book, error)
	Get(ctx context.Context) ([]entity.Book, error)
	Delete(ctx context.Context, id int) error
}
