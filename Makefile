SHELL = bash
MKDIR = mkdir

GOTEST := go test
ifneq ($(shell which gotestsum),)
	GOTEST := gotestsum --
endif

builddir:
	@mkdir -p build

cgexec: CGO_ENABLED=0
cgexec: GOOS=linux
cgexec: GOARCH=amd64
cgexec: BUILDFLAGS=-buildmode pie -tags 'osusergo netgo static_build'
cgexec: dep builddir cmd/cgexec/cgexec.go
	go build -o build/cgexec cmd/cgexec/cgexec.go
.PHONY: cgexec

clean:
	$(RM) -r build
.PHONY: clean

test: vet
	@$(GOTEST) -v -race ./...
.PHONY: test

vet: dep
	@go vet ./...
.PHONY: vet

dep:
	@go mod download
.PHONY: dep