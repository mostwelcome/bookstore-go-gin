# Bookstore API with Go and Gin

This project is a simple RESTful API for a bookstore. It is built with Go and the Gin Web Framework. It uses SQLite as the database.

## Features

- Get a list of all books
- Get a book by ID
- Create a new book
- Update an existing book
- Delete a book

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go version 1.17 or later
- SQLite3
- Gin Web Framework

### Installation

1. Clone the repo:
    ```sh
    git clone https://github.com/mostwelcome/bookstore-go-gin.git
    ```
2. Go to the project directory:
    ```sh
    cd bookstore-go-gin
    ```
3. Install the required packages:
    ```sh
    go get -d -v
    ```
4. Run the application:
    ```sh
    go run main.go
    ```

## Running the tests

You can run the tests with the following command:
```sh
go test -v
