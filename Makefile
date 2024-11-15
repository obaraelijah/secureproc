SHELL = bash
MKDIR = mkdir -p
BUILDDIR = build
COVERAGEDIR=$(BUILDDIR)/coverage
EXECUTABLES =
EXECUTABLES += $(BUILDDIR)/cgexec
EXECUTABLES += $(BUILDDIR)/jobmanager
EXECUTABLES += $(BUILDDIR)/jobctl

GOTEST := go test
ifneq ($(shell which gotestsum),)
	GOTEST := gotestsum -- 
endif

all: proto $(EXECUTABLES)

$(BUILDDIR):
	$(MKDIR) $(BUILDDIR)

$(BUILDDIR)/cgexec: CGO_ENABLED=0
$(BUILDDIR)/cgexec: GOOS=linux
$(BUILDDIR)/cgexec: GOARCH=amd64
$(BUILDDIR)/cgexec: BUILDFLAGS=-buildmode pie -tags 'osusergo netgo static_build'
$(BUILDDIR)/cgexec: dep $(BUILDDIR) cmd/cgexec/cgexec.go
	go build -race -o $(BUILDDIR)/cgexec cmd/cgexec/cgexec.go

$(BUILDDIR)/jobmanager: dep $(BUILDDIR) cmd/server/server.go
	go build -race -o $(BUILDDIR)/jobmanager cmd/server/server.go

$(BUILDDIR)/jobctl: dep $(BUILDDIR) cmd/client/client.go
	go build -race -o $(BUILDDIR)/jobctl cmd/client/client.go

proto:
	@if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	protoc --proto_path=./service/jobmanager/jobmanagerv1/ --go_out=./service/jobmanager/jobmanagerv1 --go_opt=paths=source_relative --go-grpc_out=./service/jobmanager/jobmanagerv1 --go-grpc_opt=paths=source_relative ./service/jobmanager/jobmanagerv1/jobmanager.proto
.PHONY: proto

clean:
	$(RM) -r $(BUILDDIR) ./service/v1/jobmanager_grpc.pb.go  ./service/v1/jobmanager.pb.go
.PHONY: clean

$(COVERAGEDIR):
	$(MKDIR) -p $(COVERAGEDIR)

test: vet $(COVERAGEDIR)
	@$(GOTEST) -v -race -coverprofile=${COVERAGEDIR}/coverage.out -coverpkg=./... ./...
	@go tool cover -func=${COVERAGEDIR}/coverage.out -o ${COVERAGEDIR}/function-coverage.txt
	@go tool cover -html=${COVERAGEDIR}/coverage.out -o ${COVERAGEDIR}/coverage.html
.PHONY: test

# Not using $(GOTEST) here since root might not have it installed
inttest: CERTDIR=$(shell readlink -f certs)
inttest: vet $(BUILDDIR)/cgexec
	@cp $(BUILDDIR)/cgexec /tmp
	@sudo bash -c 'CERTDIR=$(CERTDIR) go test -v -race -count=1 --tags=integration ./test/...'
.PHONY: inttest

vet: dep
	@go vet -race ./...
.PHONY: vet

dep:
	@go mod download
	@go mod tidy
.PHONY: dep