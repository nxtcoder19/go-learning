package main

import (
	"fmt"
)

type Book struct {
	Id     string
	Title  string
	Author string
}

type LibraryManagement interface {
	listBook() []Book
	getBook(id string) (Book, bool)
	addBook(book Book) Book
	deleteBook(id string) bool
}

type InMemoryBook struct {
	books map[string]Book
}

func NewInMemoryBook() LibraryManagement {
	return &InMemoryBook{
		books: make(map[string]Book),
	}
}

func (i *InMemoryBook) addBook(book Book) Book {
	i.books[book.Id] = book
	return book
}

func (i *InMemoryBook) listBook() []Book {
	var bookList []Book
	for _, book := range i.books {
		bookList = append(bookList, book)
	}
	return bookList
}

func (i *InMemoryBook) getBook(id string) (Book, bool) {
	book, exists := i.books[id]
	return book, exists
}

func (i *InMemoryBook) deleteBook(id string) bool {
	_, exists := i.books[id]
	if exists {
		delete(i.books, id)
	}
	return exists
}

func main() {
	library := NewInMemoryBook()

	book1 := Book{Id: "1", Title: "The Go Programming Language", Author: "Alan Donovan"}
	book2 := Book{Id: "2", Title: "Clean Code", Author: "Robert C. Martin"}

	library.addBook(book1)
	library.addBook(book2)

	fmt.Println("Books in Library:")
	for _, book := range library.listBook() {
		fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.Id, book.Title, book.Author)
	}

	id := "1"
	book, found := library.getBook(id)
	if found {
		fmt.Printf("\nBook found (ID: %s): Title: %s, Author: %s\n", book.Id, book.Title, book.Author)
	} else {
		fmt.Printf("\nBook with ID %s not found\n", id)
	}

	deleted := library.deleteBook("2")
	if deleted {
		fmt.Println("\nBook with ID 2 deleted successfully")
	} else {
		fmt.Println("\nBook with ID 2 not found")
	}

	fmt.Println("\nUpdated Library:")
	for _, book := range library.listBook() {
		fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.Id, book.Title, book.Author)
	}
}
