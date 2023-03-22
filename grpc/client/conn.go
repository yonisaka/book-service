package client

import (
	"flag"
	"fmt"

	"github.com/yonisaka/book-service/config"
	"github.com/yonisaka/book-service/grpc/interceptor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewGRPCConn is a constructor
func NewGRPCConn(config *config.Config) (*grpc.ClientConn, error) {
	flag.Parse()

	addr := flag.String("addr", fmt.Sprintf("%s:%s", config.GRPCService.Host, config.GRPCService.Port), "The address to connect")
	conn, err := grpc.Dial(*addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(interceptor.UnaryAuthClientInterceptor()),
		grpc.WithStreamInterceptor(interceptor.StreamAuthClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// NewGRPCConnUserService is a constructor
func NewGRPCConnUserService(config *config.Config) (*grpc.ClientConn, error) {
	flag.Parse()

	addr := flag.String("addrUserService", fmt.Sprintf("%s:%s", config.GRPCUserService.Host, config.GRPCUserService.Port), "The address to connect")
	conn, err := grpc.Dial(*addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(interceptor.UnaryAuthClientInterceptor()),
		grpc.WithStreamInterceptor(interceptor.StreamAuthClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
