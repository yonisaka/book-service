package client

import (
	"context"
	pbBook "github.com/yonisaka/protobank/book"
)

// GetBookList is a method
func (r GRPCClient) GetBookList(ctx context.Context, payload *pbBook.BookListQuery) (*pbBook.BooksResponse, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	books, err := r.book.GetBookList(ctx, &pbBook.BookListQuery{})
	if err != nil {
		return nil, err
	}

	return books, nil
}
