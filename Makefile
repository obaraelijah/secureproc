.PHONY: all client server test clean deps lint proto

# Build settings
BINARY_DIR=bin
CLIENT_BINARY=$(BINARY_DIR)/jobctl
SERVER_BINARY=$(BINARY_DIR)/secureprocd
GO=go
GOFLAGS=-trimpath
LDFLAGS=-s -w

all: deps client server

$(BINARY_DIR):
	mkdir -p $(BINARY_DIR)

client: $(BINARY_DIR)
	$(GO) build $(GOFLAGS) -ldflags "$(LDFLAGS)" -o $(CLIENT_BINARY) ./cmd/jobctl

server: $(BINARY_DIR)
	$(GO) build $(GOFLAGS) -ldflags "$(LDFLAGS)" -o $(SERVER_BINARY) ./cmd/secureprocd

test:
	$(GO) test -v ./...

clean:
	rm -rf $(BINARY_DIR)
	$(GO) clean

deps:
	$(GO) mod download
	$(GO) mod tidy

lint:
	golangci-lint run

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		pkg/api/jobmanager.proto