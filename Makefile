# Go parameters
GO := go
GOFMT := gofmt
BINARY_NAME := build/learnium
BINARY_FILE := ./build/learnium
MAIN_FILE := ./cmd/main.go

.PHONY: all
all: clean fmt binary-preview

.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)

.PHONY: fmt
fmt:
	@echo "Formatting code..."
	@$(GOFMT) -w .

.PHONY: install
install:
	go get ./...


.PHONY: build
build: clean
	@echo "Building binary..."
	@$(GO) build -o $(BINARY_NAME) $(MAIN_FILE)

.PHONY: run
run: start-storage
	@echo "Running application..."
	@$(GO) run $(MAIN_FILE)


.PHONY: binary-preview
binary-preview: build start-storage
	@echo "Previewing application..."
	@$(BINARY_FILE)


.PHONY: start-storage
start-storage:
	@echo "Starting database and cache containers"
	docker compose up -d database cache

.PHONY: stop-storage
stop-storage:
	@echo "Shutting down database container"
	docker compose down --remove-orphans database cache
	