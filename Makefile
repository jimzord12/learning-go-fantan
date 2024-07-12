build:
	@go build -o bin/myApp.exe ./cmd

test:
	@go test -v ./...

run: build
	@./bin/myApp.exe