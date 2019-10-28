package api

import "github.com/kataras/iris/v12"

// Book type with Name, Author and ISBN
type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description,omitempty"`
}

var books = map[string]Book{
	"0345391802": {Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", ISBN: "0345391802"},
	"0000000000": {Title: "Cloud Native Go", Author: "M.-Leander Reimer", ISBN: "0000000000"},
}

// AllBooks returns a slice of all books
func AllBooks() []Book {
	values := make([]Book, len(books))
	idx := 0
	for _, book := range books {
		values[idx] = book
		idx++
	}
	return values
}

// AllBooksHandler to be used as Handler for Book API
func AllBooksHandler(ctx iris.Context) {
	ctx.JSON(AllBooks())
}

// CreateBookHandler to be used as Handler for Book API
func CreateBookHandler(ctx iris.Context) {
	book := Book{}
	if err := ctx.ReadJSON(book); err != nil {
		ctx.StatusCode(400)
		return
	}
	isbn, created := CreateBook(book)

	if created {
		ctx.Redirect("/api/books/"+isbn, iris.StatusCreated)
	} else {
		ctx.StatusCode(iris.StatusConflict)
	}
}

// GetBookHandler to be used as Handler for Book API
func GetBookHandler(ctx iris.Context) {
	isbn := ctx.Params().Get("isbn")

	book, found := GetBook(isbn)
	if !found {
		ctx.StatusCode(iris.StatusNotFound)
		return

	}

	ctx.JSON(book)
}

// UpdateBookHandler to be used as Handler for Book API
func UpdateBookHandler(ctx iris.Context) {
	isbn := ctx.Params().Get("isbn")

	book := Book{ISBN: isbn}
	if err := ctx.ReadJSON(book); err != nil {
		ctx.StatusCode(400)
		return
	}

	exists := UpdateBook(isbn, book)
	if !exists {
		ctx.StatusCode(iris.StatusNotFound)
	}
}

// DeleteBookHandler to be used as Handler for Book API
func DeleteBookHandler(ctx iris.Context) {
	isbn := ctx.Params().Get("isbn")
	DeleteBook(isbn)
}

// GetBook returns the book for a given ISBN
func GetBook(isbn string) (Book, bool) {
	book, found := books[isbn]
	return book, found
}

// CreateBook creates a new Book if it does not exist
func CreateBook(book Book) (string, bool) {
	_, exists := books[book.ISBN]
	if exists {
		return "", false
	}
	books[book.ISBN] = book
	return book.ISBN, true
}

// UpdateBook updates an existing book
func UpdateBook(isbn string, book Book) bool {
	_, exists := books[isbn]
	if exists {
		books[isbn] = book
	}
	return exists
}

// DeleteBook removes a book from the map by ISBN key
func DeleteBook(isbn string) {
	delete(books, isbn)
}
