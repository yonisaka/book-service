package server

import (
	"github.com/yonisaka/book-service/config"
	"github.com/yonisaka/book-service/domain/service"
	"github.com/yonisaka/book-service/grpc/handler"
	"github.com/yonisaka/book-service/grpc/interceptor"
	pbBook "github.com/yonisaka/protobank/book"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server is struct to hold any dependencies used for server
type Server struct {
	config *config.Config
	repo   *service.Repositories
}

// NewGRPCServer is constructor
func NewGRPCServer(conf *config.Config, repo *service.Repositories) *Server {
	return &Server{
		config: conf,
		repo:   repo,
	}
}

// Run is a method gRPC server
func (s *Server) Run(port int) error {
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.UnaryLoggerServerInterceptor(),
			interceptor.UnaryAuthServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			interceptor.StreamLoggerServerInterceptor(),
			interceptor.StreamAuthServerInterceptor(),
		),
	)

	handlers := handler.NewHandler(s.config, s.repo)

	// register service server
	pbBook.RegisterBookServiceServer(server, handlers)

	// register reflection
	reflection.Register(server)

	return RunGRPCServer(server, port)
}
