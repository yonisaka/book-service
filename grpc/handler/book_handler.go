package handler

import (
	"context"

	"github.com/yonisaka/book-service/domain/entity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/yonisaka/protobank/book"
)

func (c *Handler) GetBook(ctx context.Context, r *pb.BookByIDRequest) (*pb.BookResponse, error) {
	book, err := c.repo.Book.Find(ctx, int(r.GetId()))

	if err != nil {
		return nil, status.Error(codes.NotFound, "Data Not Found")
	}

	return &pb.BookResponse{
		Id:          uint64(book.ID),
		Title:       book.Title,
		Description: book.Description,
		Author:      book.Author,
		CreatedAt:   book.CreatedAt.String(),
		UpdatedAt:   book.UpdatedAt.String(),
	}, nil
}

func (c *Handler) GetBookList(ctx context.Context, r *pb.BookListQuery) (*pb.BooksResponse, error) {
	serv, err := c.repo.Book.Get(ctx)

	if err != nil {
		return nil, err
	}

	var books []*pb.BookResponse
	for _, b := range serv {
		books = append(books, &pb.BookResponse{
			Id:          uint64(b.ID),
			Title:       b.Title,
			Description: b.Description,
			Author:      b.Author,
			CreatedAt:   b.CreatedAt.String(),
			UpdatedAt:   b.UpdatedAt.String(),
		})
	}

	return &pb.BooksResponse{
		Books: books,
	}, nil
}

func (c *Handler) CreateBook(ctx context.Context, r *pb.BookCreateRequest) (*pb.BookResponse, error) {
	book := entity.Book{
		Title:       r.GetTitle(),
		Description: r.GetDescription(),
		Author:      r.GetAuthor(),
	}

	err := c.repo.Book.Create(ctx, &book)

	if err != nil {
		return nil, err
	}

	return &pb.BookResponse{
		Id:          uint64(book.ID),
		Title:       book.Title,
		Description: book.Description,
		Author:      book.Author,
		CreatedAt:   book.CreatedAt.String(),
		UpdatedAt:   book.UpdatedAt.String(),
	}, nil
}

func (c *Handler) UpdateBook(ctx context.Context, r *pb.BookUpdateRequest) (*pb.BookResponse, error) {
	bookId := int(r.GetId())

	if _, err := c.repo.Book.Find(ctx, bookId); err != nil {
		return nil, status.Error(codes.NotFound, "Data Not Found")
	}

	bookData := &entity.Book{
		Title:       r.GetTitle(),
		Description: r.GetDescription(),
		Author:      r.GetAuthor(),
	}

	err := c.repo.Book.Update(ctx, bookId, bookData)
	if err != nil {
		return nil, err
	}

	return &pb.BookResponse{
		Id:          uint64(bookId),
		Title:       bookData.Title,
		Description: bookData.Description,
		Author:      bookData.Author,
		CreatedAt:   bookData.CreatedAt.String(),
		UpdatedAt:   bookData.UpdatedAt.String(),
	}, nil
}

func (c *Handler) DeleteBook(ctx context.Context, r *pb.BookByIDRequest) (*pb.BookDeleteResponse, error) {
	bookId := int(r.GetId())

	if _, err := c.repo.Book.Find(ctx, bookId); err != nil {
		return nil, status.Error(codes.NotFound, "Data not found")
	}

	err := c.repo.Book.Delete(ctx, int(r.GetId()))

	if err != nil {
		return nil, err
	}

	return &pb.BookDeleteResponse{
		Message: "ok",
	}, nil
}
