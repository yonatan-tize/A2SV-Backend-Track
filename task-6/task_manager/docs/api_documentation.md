# Task Manager Project

## Overview

The Task Manager project is a simple task management system built with Go. It allows users to create, read, update, and delete tasks. The project uses the Gin framework for the web server and the official MongoDB driver for database interactions.

## Setup Instructions

### Prerequisites

- Go (version 1.16 or higher)
- A running MongoDB instance

### Installation


1. **Set up the database:**

    Update the database connection string in `data/task_service.go` with your own MongoDB connection parameters.

    ```go
    // Example connection string
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    ```

2. **Install dependencies:**

    ```sh
    go mod tidy
    ```

3. **Run the application:**

    ```sh
    go run main.go
    ```

## API Documentation

You can refer to the detailed API documentation using the link below:

[Postman API Documentation](https://documenter.getpostman.com/view/37171778/2sA3s1orgE)

### Endpoints

