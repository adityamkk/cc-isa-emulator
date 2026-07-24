BINARY_NAME := cc-isa-emulator
BUILD_DIR := bin

.PHONY: run build clean test vet fmt

build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) .

run:
	go run . --bin=README.md --cores=15

clean:
	rm -rf $(BUILD_DIR)

test:
	go test ./...

vet:
	go vet ./...
