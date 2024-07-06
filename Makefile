build:
	@go build -o bin/myApp ./cmd

test:
	@go test -v ./...

run: build
	@./bin/myApp