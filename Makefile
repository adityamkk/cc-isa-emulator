BINARY_NAME := cc-isa-emulator
BUILD_DIR := bin

.PHONY: run build clean test vet fmt

build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) .

run:
	go run . $(ARGS)

clean:
	rm -rf $(BUILD_DIR)

test:
	go test ./...

vet:
	go vet ./...
