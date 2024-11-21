# Find all .go files in the directory and its subdirectories
GO_FILES := $(shell find . -name '*.go')

# Target for building the shiftsync binary
build: $(GO_FILES)
	@go build -o ./shiftsync ./cmd/shiftsync
