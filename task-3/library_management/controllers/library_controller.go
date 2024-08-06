package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"library_management/models"
	"library_management/services"
)

var library = services.Library{
	Books: make(map[int]models.Book),
	Members: map[int]models.Member{
		1: {ID: 1, Name: "John", BorrowedBooks: make([]models.Book, 0)},
		2: {ID: 2, Name: "James", BorrowedBooks: make([]models.Book, 0)},
		3: {ID: 3, Name: "Jack", BorrowedBooks: make([]models.Book, 0)},
	},
}

func MainProgram() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Library Management System")
		fmt.Println("Enter 1 to add a new book")
		fmt.Println("Enter 2 to remove an existing book")
		fmt.Println("Enter 3 to borrow a book")
		fmt.Println("Enter 4 to return a book")
		fmt.Println("Enter 5 to list all available books")
		fmt.Println("Enter 6 to list all borrowed books by a member")
		fmt.Println("Enter 7 to exit")
		fmt.Print("Enter your choice: ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		switch userInput {
		case "1":
			addBook(&library)
		case "2":
			removeBook(&library)
		case "3":
			borrowBook(&library)
		case "4":
			returnBook(&library)
		case "5":
			listAvailableBooks(&library)
		case "6":
			listBorrowedBooks(&library)
		case "7":
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

// addBook adds a book to the library.
// It takes a pointer to a services.Library as a parameter.
func addBook(library *services.Library) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter ID of the book: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)
	bookID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Invalid book ID.")
		return
	}

	fmt.Print("Enter title of the book: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Printf("Enter author of the book: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	newBook := models.Book{
		ID:     bookID,
		Title:  title,
		Author: author,
		Status: "Available",
	}
	library.AddBook(newBook)
	fmt.Println("Book added successfully.")
}

// removeBook removes a book from the library.
// It takes a pointer to a services.Library as a parameter.
func removeBook(library *services.Library) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Book ID to remove: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)
	bookID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Invalid book ID.")
		return
	}

	library.RemoveBook(bookID)
	fmt.Println("Book removed successfully.")
}

// borrowBook is a function that allows a user to borrow a book from the library.
// It takes a pointer to a services.Library object as a parameter.
func borrowBook(library *services.Library) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter book ID to borrow: ")
	bookIDStr, _ := reader.ReadString('\n')
	bookIDStr = strings.TrimSpace(bookIDStr)
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		fmt.Println("Invalid book ID.")
		return
	}

	fmt.Print("Enter Member ID: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberIDStr = strings.TrimSpace(memberIDStr)
	memberID, err := strconv.Atoi(memberIDStr)
	if err != nil {
		fmt.Println("Invalid member ID.")
		return
	}

	err = library.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book borrowed successfully.")
	}
}

// returnBook returns a book to the library.
// It takes a pointer to a services.Library object as a parameter.
func returnBook(library *services.Library) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Book ID to return: ")
	bookIDStr, _ := reader.ReadString('\n')
	bookIDStr = strings.TrimSpace(bookIDStr)
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		fmt.Println("Invalid book ID.")
		return
	}

	fmt.Print("Enter Member ID: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberIDStr = strings.TrimSpace(memberIDStr)
	memberID, err := strconv.Atoi(memberIDStr)
	if err != nil {
		fmt.Println("Invalid member ID.")
		return
	}

	err = library.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book returned successfully.")
	}
}
// listAvailableBooks prints the list of available books in the library.
// It takes a pointer to a services.Library object as a parameter.
// It retrieves the list of available books using the ListAvailableBooks method of the library object.
// Then, it prints the ID, title, and author of each book in the list.
func listAvailableBooks(library *services.Library) {
	books := library.ListAvailableBooks()
	fmt.Println("Available Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

// listBorrowedBooks displays a list of borrowed books for a given member ID.
// It prompts the user to enter a member ID, retrieves the list of borrowed books from the library service,
// and prints the details of each book.
//
// Parameters:
// - library: A pointer to the services.Library instance.
//
// Example usage:
//   library := &services.Library{}
//   listBorrowedBooks(library)
func listBorrowedBooks(library *services.Library) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Member ID: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberIDStr = strings.TrimSpace(memberIDStr)
	memberID, err := strconv.Atoi(memberIDStr)
	if err != nil {
		fmt.Println("Invalid member ID.")
		return
	}

	books, err:= library.ListBorrowedBooks(memberID)
	if books != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Borrowed Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}
