GO		:= $(shell which go)
GOFMT		:= $(shell which gofmt)
GOPATH		:= $(shell go env GOPATH)
GOBIN		:= $(GOPATH)/bin
GOLINT		:= $(GOBIN)/golint
GOSEC		:= $(GOBIN)/gosec

COVERAGE_FILE	:= coverage.out

.PHONY: build
build:  clean format lint vet sec coverage seed.sql music-api

.PHONY: clean
clean:
	rm -f seed.sql music-api

.PHONY: sec-install
sec-install:
	$(GO) install github.com/securego/gosec/v2/cmd/gosec@latest

.PHONY: lint-install
lint-install:
	test -e $(GOLINT) || $(GO) install golang.org/x/lint/golint@latest

.PHONY: format
format:
	@./check-formatting.sh

.PHONY: lint
lint: lint-install
	$(GOLINT) -set_exit_status ./...

.PHONY: sec
sec: sec-install
	$(GOSEC) ./...

.PHONY: vet
vet:
	$(GO) vet ./...

.PHONY: test
test:
	$(GO) test -race -v ./... -coverprofile=$(COVERAGE_FILE)

.PHONY: coverage
coverage: test
	$(GO) tool cover -func=$(COVERAGE_FILE)

.PHONY: coverage-html
coverage-html: test
	$(GO) tool cover -html=$(COVERAGE_FILE)

seed.sql:
	$(GO) run ./cmd/gen -output=$(CURDIR)/seed.sql

music-api:
	$(GO) build -o music-api ./cmd/api
