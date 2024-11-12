SHELL = bash
MKDIR = mkdir -p
BUILDDIR = build
COVERAGEDIR=$(BUILDDIR)/coverage
EXECUTABLES += $(BUILDDIR)/test-blkiolimit
EXECUTABLES += $(BUILDDIR)/test-memorylimit

GOTEST := go test
ifneq ($(shell which gotestsum),)
	GOTEST := gotestsum -- 
endif

all: $(BUILDDIR) $(BUILDDIR)/cgexec $(EXECUTABLES)

$(BUILDDIR):
	$(MKDIR) $(BUILDDIR)

$(BUILDDIR)/cgexec: CGO_ENABLED=0
$(BUILDDIR)/cgexec: GOOS=linux
$(BUILDDIR)/cgexec: GOARCH=amd64
$(BUILDDIR)/cgexec: BUILDFLAGS=-buildmode pie -tags 'osusergo netgo static_build'
$(BUILDDIR)/cgexec: dep $(BUILDDIR) cmd/cgexec/cgexec.go
	go build -race -o $(BUILDDIR)/cgexec cmd/cgexec/cgexec.go

$(BUILDDIR)/test-blkiolimit: dep $(BUILDDIR) test/job/blkiolimit/blkiolimit.go
	go build -race -o $(BUILDDIR)/test-blkiolimit test/job/blkiolimit/blkiolimit.go

$(BUILDDIR)/test-memorylimit: dep $(BUILDDIR) test/job/memorylimit/memorylimit.go
	go build -race -o $(BUILDDIR)/test-memorylimit test/job/memorylimit/memorylimit.go

clean:
	$(RM) -r $(BUILDDIR)
.PHONY: clean

$(COVERAGEDIR):
	$(MKDIR) -p $(COVERAGEDIR)

test: vet $(COVERAGEDIR)
	@$(GOTEST) -v -race -coverprofile=${COVERAGEDIR}/coverage.out -coverpkg=./... ./...
	@go tool cover -func=${COVERAGEDIR}/coverage.out -o ${COVERAGEDIR}/function-coverage.txt
	@go tool cover -html=${COVERAGEDIR}/coverage.out -o ${COVERAGEDIR}/coverage.html
.PHONY: test

# Not using $(GOTEST) here since root might not have it installed
inttest: vet $(BUILDDIR)/cgexec
	@cp $(BUILDDIR)/cgexec /tmp
	@sudo go test -v -race --tags=integration ./test/...

vet: dep
	@go vet -race ./...
.PHONY: vet

dep:
	@go mod download
.PHONY: dep