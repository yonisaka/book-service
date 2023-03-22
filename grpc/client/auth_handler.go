package client

import (
	"context"

	pbAuth "github.com/yonisaka/protobank/auth"
)

// AuthB2B is a method
func (r GRPCClient) AuthB2B(ctx context.Context, payload *pbAuth.AuthB2BPayload) (*pbAuth.UserResponse, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	user, err := r.auth.AuthB2B(ctx, payload)
	if err != nil {
		return nil, err
	}

	return user, nil
}
