// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main

import (
	"fmt"
	"sync"
)

type Book struct {
	Id     string
	Title  string
	Author string
}

type LibraryManagement interface {
	storeBook(book Book) Book
	getBook(id string) *Book
}

type InMemoryBook struct {
	books map[string]Book
	mu    sync.RWMutex
}

func NewInMemoryBook() LibraryManagement {
	return &InMemoryBook{
		books: make(map[string]Book),
	}
}

func (i *InMemoryBook) storeBook(book Book) Book {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.books[book.Id] = book
	return book
}

func (i *InMemoryBook) getBook(id string) *Book {
	i.mu.RLock()
	defer i.mu.RUnlock()
	book, exists := i.books[id]
	if !exists {
		return nil
	}

	return &book
}

func main() {
	fmt.Println("Try programiz.pro")
	library := NewInMemoryBook()

	var wg sync.WaitGroup

	wg.Add(2)

	book1 := Book{
		Id:     "1",
		Title:  "Go Program",
		Author: "sam",
	}

	book2 := Book{
		Id:     "2",
		Title:  "Python Program",
		Author: "john",
	}

	go func() {
		defer wg.Done()
		library.storeBook(book1)
	}()

	go func() {
		defer wg.Done()
		library.storeBook(book2)
	}()

	wg.Wait()

	//   library.storeBook(book1)
	//   library.storeBook(book2)

	wg.Add(2)
	id1 := "1"
	id2 := "2"

	go func() {
		defer wg.Done()
		book := library.getBook(id1)
		if book == nil {
			fmt.Println("no book found")
		} else {
			fmt.Printf("Id: %d, Title: %s, Author: %s\n", book.Id, book.Title, book.Author)
		}
	}()

	go func() {
		defer wg.Done()
		book := library.getBook(id2)
		if book == nil {
			fmt.Println("no book found")
		} else {
			fmt.Printf("Id: %d, Title: %s, Author: %s\n", book.Id, book.Title, book.Author)
		}
	}()

	//   id1 := "1"
	//   id2 := "2"
	//   book := library.getBook(id)
	//   if book == nil {
	//       fmt.Println("no book found")
	//   } else {
	//       fmt.Printf("Id: %d, Title: %s, Author: %s\n", book.Id, book.Title, book,Author)
	//   }

	wg.Wait()

	fmt.Println("All Done")

}

// Make a library system where users can borrow and place books in a bookshelves. Make use of map structure to store [books]=bookselves (ie. map[string]int). Make sure the users can store or search books concurrently.
