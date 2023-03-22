package contract

// ProtectedMethods is a function to hold grpc service methods
// false value indicates that the method is not protected (no authorization needed)
func ProtectedMethods() map[string]bool {
	return map[string]bool{
		"/log.LogService/SaveHttpLog":   true,
		"/book.BookService/GetBookList": true,
		"/book.BookService/GetBook":     true,
		"/book.BookService/CreateBook":  true,
		"/book.BookService/UpdateBook":  true,
		"/book.BookService/DeleteBook":  true,
	}
}
