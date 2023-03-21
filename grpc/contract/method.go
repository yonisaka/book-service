package contract

// ProtectedMethods is a function to hold grpc service methods
// false value indicates that the method is not protected (no authorization needed)
func ProtectedMethods() map[string]bool {
	return map[string]bool{
		"/book.BookService/GetBookList": true,
	}
}
