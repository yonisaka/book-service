package repository

import (
	"context"

	"github.com/yonisaka/book-service/domain/entity"
)

// HttpLogRepositoryInterface is contract
type HttpLogRepositoryInterface interface {
	Get(ctx context.Context) ([]*entity.HttpLog, error)
	Find(ctx context.Context, id int) (*entity.HttpLog, error)
	Save(ctx context.Context, log *entity.HttpLog) error
}
