package handler

import (
	"github.com/yonisaka/book-service/domain/service"
	"github.com/yonisaka/book-service/grpc/client"
)

// Handler is a struct
type Handler struct {
	client *client.GRPCClient
	repo   *service.Repositories
}

// NewHandler is a function
func NewHandler(
	repo *service.Repositories,
	client *client.GRPCClient,
) *Handler {
	return &Handler{
		repo:   repo,
		client: client,
	}
}
