# Library Management System Documentation

## Overview
This library management system is a simple command-line application written in Go. It allows users to add, remove, borrow, and return books. It also enables listing available and borrowed books.

## Structure
- `main.go`: The entry point of the application.
- `controllers/`: Contains the controller logic for handling user input.
  - `library_controller.go`: Manages the main program logic and user interactions.
- `models/`: Contains the data models for the application.
  - `book.go`: Defines the `Book` struct.
  - `member.go`: Defines the `Member` struct.
- `services/`: Contains the business logic and core functionality of the application.
  - `library_service.go`: Implements the `Library` struct and `LibraryManager` interface.
- `docs/`: Contains documentation and additional resources.

## Usage
1. **Add a New Book**: Allows the user to add a new book to the library.
2. **Remove an Existing Book**: Allows the user to remove a book from the library by its ID.
3. **Borrow a Book**: Allows a member to borrow a book by its ID.
4. **Return a Book**: Allows a member to return a borrowed book by its ID.
5. **List All Available Books**: Displays all books that are available for borrowing.
6. **List All Borrowed Books by a Member**: Displays all books borrowed by a specific member.

## Running the Application
1. Navigate to the project directory.
2. Run the application using `go run main.go`.


