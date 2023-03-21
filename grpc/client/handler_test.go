package client_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yonisaka/book-service/grpc/client"
)

func TestRun(t *testing.T) {
	err := client.Run()

	assert.NoError(t, err)
}