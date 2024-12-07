# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=main
BINARY_UNIX=$(BINARY_NAME)_unix

# Build the project
.PHONY: build
build:
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/txparser

# Run the tests
.PHONY: test
test:
	$(GOTEST) -v ./...

# Clean the build files
.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# Run the application
.PHONY: run
run:
	$(GOCMD) run ./cmd/txparser

# Build and run Docker container
.PHONY: docker
docker:
	docker build -t ethereum-parser .
	docker run -p 8080:8080 ethereum-parser

# Format the code
.PHONY: fmt
fmt:
	$(GOCMD) fmt ./...

# Format the code using gofmt
.PHONY: gofmt
gofmt:
	gofmt -w .
