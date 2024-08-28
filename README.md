## Bookstore Management API

A project to build a RESTful API for managing a bookstore using Go (Golang). The project covers basic concepts such as database integration, JSON handling, and server routing.

### Table of Contents

1. [Introduction](#introduction)
2. [Technologies Used](#technologies-used)
3. [Project Structure](#project-structure)
4. [Getting Started](#getting-started)
5. [Usage](#usage)
6. [Contributing](#contributing)

### Introduction

This project is designed to help beginners understand how to create a basic API using Go. It covers essential topics such as:

- Database connection using MySQL
- ORM with GORM
- JSON Marshalling and Unmarshalling
- Project structure and best practices
- Routing with Gorilla Mux

### Technologies Used

- **Go (Golang)**
- **MySQL**
- **GORM** - Go ORM library for interacting with databases
- **Gorilla Mux** - HTTP request router and dispatcher for Go

### Project Structure

```
go-bookstore/
│
├── cmd/
│   └── main/
│       └── main.go            # Main application entry point
│
├── pkg/
│   ├── config/
│   │   └── app.go             # Application configuration and setup
│   ├── controllers/
│   │   └── book-controller.go # Controllers for handling book-related requests
│   ├── models/
│   │   └── book.go            # Book model and database schema
│   ├── routes/
│   │   └── bookstore-routes.go# Application routes
│   └── utils/
│       └── utils.go           # Utility functions and helpers
│
├── go.mod                     # Go module file
├── go.sum                     # Go dependencies checksum
└── README.md                  # Project README
```

### Getting Started

To get a local copy up and running, follow these steps:

1. **Clone the repository**
   ```bash
   git clone [https://github.com/abdullahalhwyji/go-bookstore.git]
   cd go-bookstore
   ```

2. **Install Go and MySQL**  
   Make sure you have Go and MySQL installed on your machine.

3. **Install dependencies**
   ```bash
   go mod tidy
   ```

4. **Set up the database**
   - Create a MySQL database for the project.
   - Update the database configuration in `pkg/config/app.go`.

5. **Run the application**
   ```bash
   go run cmd/main/main.go
   ```

### Usage

- **CRUD Operations**: The API supports standard CRUD operations for managing bookstore items.
- **Routing**: All routes are defined in `pkg/routes/bookstore-routes.go`.
- **Controllers**: Business logic for each route is handled in `pkg/controllers/book-controller.go`.

### Contributing

Contributions are welcome! Please fork the repository and create a pull request with your changes. 

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/YourFeature`)
3. Commit your Changes (`git commit -m 'Add YourFeature'`)
4. Push to the Branch (`git push origin feature/YourFeature`)
5. Open a Pull Request

