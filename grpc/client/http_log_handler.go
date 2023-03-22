package client

import (
	"context"
	"fmt"
	pbLog "github.com/yonisaka/protobank/log"
)

// SaveHttpLog is a method
func (r GRPCClient) SaveHttpLog(ctx context.Context, payload *pbLog.SaveHttpLogRequest) (*pbLog.HttpLog, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	rsp, err := r.httpLog.SaveHttpLog(ctx, &pbLog.SaveHttpLogRequest{
		Ip:     payload.Ip,
		Path:   payload.Path,
		Method: payload.Method,
	})
	if err != nil {
		return nil, fmt.Errorf("create stream, %w", err)
	}
	return rsp, nil
}
