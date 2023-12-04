BINARY_NAME=day01

.PHONY: build run clean test test_coverage dep vet lint

all: clean build

build:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin cmd/${BINARY_NAME}/main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux cmd/${BINARY_NAME}/main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows cmd/${BINARY_NAME}/main.go

# Run the application
run:
	@echo "Running..."
	./${BINARY_NAME}

clean:
	@echo "Cleaning..."
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows

test:
	@echo "Testing..."
	go test ./...

test_coverage:
	@echo "Testing coverage..."
	go test ./... -coverprofile=coverage.out

dep:
	go mod download

vet:
	go vet

lint:
	golangci-lint run --enable-all
