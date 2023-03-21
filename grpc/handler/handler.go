package handler

import (
	"github.com/yonisaka/book-service/config"
	"github.com/yonisaka/book-service/domain/service"
	pbBook "github.com/yonisaka/protobank/book"
)

// Interface is an interface
type Interface interface {
	pbBook.BookServiceServer
}

// Handler is struct
type Handler struct {
	config *config.Config
	repo   *service.Repositories

	pbBook.UnimplementedBookServiceServer
}

// NewHandler is a constructor
func NewHandler(conf *config.Config, repo *service.Repositories) *Handler {
	return &Handler{
		config: conf,
		repo:   repo,
	}
}

var _ Interface = &Handler{}
