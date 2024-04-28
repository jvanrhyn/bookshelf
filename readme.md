
## **Book Library API**

## Overview

The Book Library API is a RESTful API built using Go, Fiber, Postgres SQL, and Slog for logging. The API manages a collection of books and keeps track of a user's progress with each book. Authentication is handled using Auth0.

## Features

- Create, read, update, and delete books    
- Register and login users
- Add and update progress for a user and book
- Authentication using Auth0

## Technologies

- Go 1.19
- Fiber 2.23
- Postgres SQL 14.2
- Slog 1.8
- Auth0

## Getting Started

## Prerequisites

- Go 1.19 or higher
- Postgres SQL 14.2 or higher
- Auth0 account (to-be-added)

## Installation

1. Clone the repository: `git clone (https://github.com/jvanrhyn/bookshelf)
2. Create a Postgres database and update the `database.yml` file with your credentials
3. Run `go mod download` to download dependencies
4. Run `go run main.go` to start the API

## Environment

Your application requires several environment variables to function correctly. These variables should be added to
a `.env` file in the root directory of your project.

The required environment variables are:

- `DB_USER`: The username for your Postgres database.
- `DB_PASS`: The password for your Postgres database.
- `DB_NAME`: The name of your Postgres database.
- `DB_HOST`: The host of your Postgres database.
- `DB_PORT`: The port of your Postgres database.
- `AUTH0_DOMAIN`: Your Auth0 domain.
- `AUTH0_CLIENT_ID`: Your Auth0 client ID.
- `AUTH0_CLIENT_SECRET`: Your Auth0 client secret.

Here is a sample `.env` file:

## API Endpoints

- `POST /books`: Create a new book
- `GET /books/:id`: Retrieve a book by ID
- `GET /books/isbn/:isbn`: Retrieve a book by ISBN
- `PUT /books/:id`: Update a book
- `DELETE /books/:id`: Delete a book
- `POST /users`: Register a new user
- `POST /login`: Login a user
- `GET /users/:id`: Retrieve a user's profile
- `GET /users/email/:email`: Retrieve a user's profile by email
- `POST /progress`: Add a new progress entry
- `GET /progress/:book_id`: Retrieve a user's progress for a book
- `PUT /progress/:book_id`: Update a user's progress for a book

## Authentication

- Use the `Authorization` header with a valid Auth0 token to authenticate requests

## Contributing

Contributions are welcome! Please open a pull request with your changes.

## License

MIT License

## Contact

[Johan van Rhyn](https://squarehole.dev)
