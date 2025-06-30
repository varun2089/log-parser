.PHONY: clean lint install test build

clean:
	@echo "Cleaning up..."
	rm -rf dist
	go mod tidy

lint:
	@echo "Running linter..."
	gofmt -l .
	golangci-lint run --timeout 5m

install:
	@echo "Installing dependencies..."
	go mod download

test: clean install
	@echo "Running tests..."
	go test -v ./...

build: clean install
	@echo "Building the project..."
	mkdir -p dist
	go build -o ./dist/log-parser cmd/main.go

run: install
	@echo "Running main.go..."
	go run cmd/main.go