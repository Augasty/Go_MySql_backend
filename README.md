# Go_MySql_backend

Packages used: 

1. GORM is an Object Relational Mapping (ORM) library for Golang. ORM converts data between incompatible type systems using object-oriented programming languages.

2. Gorilla Mux provides functionalities for matching routes, serving static files, serving single-page applications (SPAs), middleware, handling CORS requests, and testing handlers.





| routes    | functions   | endpoints  | methods |
|-----------|-------------|------------|---------|
| GET ALL   |  getBooks   | /books     | GET     |
| GET BY ID | getBook     | /books/id  | GET     |
| CREATE    | createBook  | /books     | POST    |
| UPDATE    | updateBook  | /books/id  | PUT     |
| DELETE    | deleteBook  | /books/id  | DELETE  |



To run it, first setup mysql with proper user, password and permissions.
Then, inside the cmd/main folder
go build
go run main.go