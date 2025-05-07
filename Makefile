run:
	go mod tidy && go run ./cmd/main.go

unit:
	go test -v ./test