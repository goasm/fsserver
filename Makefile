BUILD_FLAGS := -ldflags "-s -w"
OUTPUT_PATH := ./bin/fsserver

.PHONY: default

default: build

build:
	go build -o $(OUTPUT_PATH) $(BUILD_FLAGS) ./cmd/fsserver
	@echo "Build done"
