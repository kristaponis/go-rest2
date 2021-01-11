# Go REST API example

Just simple REST example No. 2

### Requirements

Content type is application/json

- GET /books
  - return list of books in JSON format
- GET /books/{id}
  - return book by id in JSON format
- POST /books
  - create new book
- POST /books
  - return Status 415 if content is not application/json
- GET /admin
  - access with basic auth
- GET /books/rand
  - return random book

