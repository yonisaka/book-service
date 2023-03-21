package route

import (
	"github.com/gin-gonic/gin"
	"github.com/yonisaka/book-service/config"
	"github.com/yonisaka/book-service/domain/service"
	"github.com/yonisaka/book-service/grpc/client"
	"github.com/yonisaka/book-service/rest/handler"
	"github.com/yonisaka/book-service/rest/middleware"
)

// WithConfig is function
func WithConfig(config *config.Config) RouterOption {
	return func(r *Router) {
		r.config = config
	}
}

// WithRepository is function
func WithRepository(repo *service.Repositories) RouterOption {
	return func(r *Router) {
		r.repo = repo
	}
}

// WithGRPCClient is function
func WithGRPCClient(client *client.GRPCClient) RouterOption {
	return func(r *Router) {
		r.client = client
	}
}

// Init is a function
func (r *Router) Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	e := gin.Default()
	e.Use(middleware.Logger())

	hand := handler.NewHandler(r.repo, r.client)

	book := handler.NewBookHandler(hand)

	e.GET("/api/books", book.GetBookList)

	return e
}
