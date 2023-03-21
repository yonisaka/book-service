package client

import (
	"flag"
	"fmt"

	"github.com/yonisaka/book-service/config"
	"github.com/yonisaka/book-service/grpc/interceptor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverHost = "localhost"
	serverPort = 9001
	DSN        = fmt.Sprintf("%s:%d", serverHost, serverPort)

	userServiceHost = "localhost"
	userServicePort = 9002
	userServiceDSN  = fmt.Sprintf("%s:%d", userServiceHost, userServicePort)
)

var (
	addr            = flag.String("addr", DSN, "The address to connect")
	addrUserService = flag.String("addrUserService", userServiceDSN, "The address to connect")
)

// NewGRPCConn is a constructor
func NewGRPCConn(_ *config.Config) (*grpc.ClientConn, error) {
	flag.Parse()

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
func NewGRPCConnUserService(_ *config.Config) (*grpc.ClientConn, error) {
	flag.Parse()

	conn, err := grpc.Dial(*addrUserService,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(interceptor.UnaryAuthClientInterceptor()),
		grpc.WithStreamInterceptor(interceptor.StreamAuthClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
