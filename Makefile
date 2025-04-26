APP_NAME := rago
BUILD_DIR := .

all: build
.PHONY: all

build:
	@echo "🔧 Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) main.go
	@echo "✅ Build complete: $(BUILD_DIR)/$(APP_NAME)"
.PHONY: build

run:
	go run main.go
.PHONY: run

clean:
	@echo "🧹 Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@echo "✅ Clean complete"
.PHONY: clean
