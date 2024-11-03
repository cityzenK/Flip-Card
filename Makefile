build:
	@go build -o bin/cards cmd/main.go

test: build
	@go test -v ./...

run:
	@./bin/cards
