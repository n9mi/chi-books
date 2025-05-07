# chi-books
REST API to create, delete, update, and view list of books that stored in memory.

## Endpoints
- <strong style="color: green"> GET </strong> `/api/v1/books`: To list books
- <strong style="color: green"> GET </strong> `/api/v1/books/{id}`: To get a book by ID
- <strong style="color: blue"> POST </strong> `/api/v1/books`: To create new book
- <strong style="color: yellow"> PUT </strong> `/api/v1/books/{id}`: To edit a book by ID
- <strong style="color: red"> DELETE </strong> `/api/v1/books/{id}`: To delete a book by ID

## How to run
Make sure you have Go installed in your system
```bash
go mod tidy
go run ./cmd/main.go
```
or use make
```bash
make run
```

## How to test
Test are available in `/test` and you can see the results in [this page](https://github.com/n9mi/chi-books/actions). You can run the test locally by
```bash
go mod tidy
go test -v ./test
```
or use make
```bash
make unit
```
Postman collection also available in `/postman`