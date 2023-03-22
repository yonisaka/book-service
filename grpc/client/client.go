package client

import (
	pbAuth "github.com/yonisaka/protobank/auth"
	pbBook "github.com/yonisaka/protobank/book"
	pbLog "github.com/yonisaka/protobank/log"
	"google.golang.org/grpc"
)

// GRPCClient is a struct
type GRPCClient struct {
	book    pbBook.BookServiceClient
	auth    pbAuth.AuthClient
	httpLog pbLog.LogServiceClient
}

// NewGRPCClient is constructor
func NewGRPCClient(conn grpc.ClientConnInterface, userServiceConn grpc.ClientConnInterface) *GRPCClient {
	return &GRPCClient{
		book:    pbBook.NewBookServiceClient(conn),
		auth:    pbAuth.NewAuthClient(userServiceConn),
		httpLog: pbLog.NewLogServiceClient(userServiceConn),
	}
}
