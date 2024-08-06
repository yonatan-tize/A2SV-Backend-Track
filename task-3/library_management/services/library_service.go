package services

import (
	"errors"
	"library_management/models"
)

// LibraryManager is an interface that defines the operations for managing a library.
type LibraryManager interface {
	// AddBook adds a book to the library.
	AddBook(book models.Book)

	// RemoveBook removes a book from the library.
	RemoveBook(bookID int)

	// BorrowBook borrows a book from the library for a specific member.
	// It returns an error if the book is not available or if the member has already borrowed the maximum number of books.
	BorrowBook(bookID int, memberID int) error

	// ReturnBook returns a borrowed book to the library.
	// It returns an error if the book is not borrowed by the specified member.
	ReturnBook(bookID int, memberID int) error

	// ListAvailableBooks returns a list of all available books in the library.
	ListAvailableBooks() []models.Book

	// ListBorrowedBooks returns a list of books borrowed by a specific member.
	// It returns an error if the member does not exist or if there are no borrowed books for the member.
	ListBorrowedBooks(memberID int) ([]models.Book, error)
}

type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

func (lib *Library) AddBook(book models.Book) {
	lib.Books[book.ID] = book
}

func (lib *Library) RemoveBook(bookID int) {
	delete(lib.Books, bookID)
}

func (lib *Library) BorrowBook(bookID int, memberID int) error {
	book, found := lib.Books[bookID]
	if !found {
		return errors.New("book not found")
	}

	if book.Status != "Available" {
		return errors.New("book is not available")
	}

	member, found := lib.Members[memberID]
	if !found {
		return errors.New("member not found")
	}

	book.Status = "Borrowed"
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	lib.Books[bookID] = book
	lib.Members[memberID] = member

	return nil
}

func (lib *Library) ReturnBook(bookID int, memberID int) error {
	book, found := lib.Books[bookID]
	if !found {
		return errors.New("book not found")
	}

	member, found := lib.Members[memberID]
	if !found {
		return errors.New("member not found")
	}

	for i, borrowedBook := range member.BorrowedBooks {
		if borrowedBook.ID == bookID {
			book.Status = "Available"
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			lib.Books[bookID] = book
			lib.Members[memberID] = member
			return nil
		}
	}
	return errors.New("the member did not borrow this book")
}

func (lib *Library) ListAvailableBooks() []models.Book {
	availableBooks := []models.Book{}
	for _, book := range lib.Books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (lib *Library) ListBorrowedBooks(memberID int) ([]models.Book, error) {
	member, found := lib.Members[memberID]
	if !found {
		return nil, errors.New("member not found")
	}
	return member.BorrowedBooks, nil
}
