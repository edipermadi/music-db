GO		:= $(shell which go)
GOFMT		:= $(shell which gofmt)
GOPATH		:= $(shell go env GOPATH)
GOBIN		:= $(GOPATH)/bin
GOLINT		:= $(GOBIN)/golangci-lint
GOSEC		:= $(GOBIN)/gosec

COVERAGE_FILE	:= coverage.out

.PHONY: build
build:  clean format lint vet sec coverage seed.sql music-api music-gen

.PHONY: clean
clean:
	rm -f seed.sql music-api music-gen

.PHONY: sec-install
sec-install:
	$(GO) install github.com/securego/gosec/v2/cmd/gosec@latest

.PHONY: lint-install
lint-install:
	test -e $(GOLINT) || $(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: format
format:
	@./check-formatting.sh

.PHONY: lint
lint: lint-install
	$(GOLINT) run ./...

.PHONY: sec
sec: sec-install
	$(GOSEC) ./...

.PHONY: vet
vet:
	$(GO) vet ./...

pkg/illustations/DroidSansFallback.ttf:
	wget https://github.com/jenskutilek/free-fonts/raw/master/Droid/Droid%20Sans%20Fallback/DroidSansFallback.ttf -O $@

.PHONY: test
test: pkg/illustations/DroidSansFallback.ttf
	$(GO) test -race -v ./... -coverprofile=$(COVERAGE_FILE)

.PHONY: coverage
coverage: test
	$(GO) tool cover -func=$(COVERAGE_FILE)

.PHONY: coverage-html
coverage-html: test
	$(GO) tool cover -html=$(COVERAGE_FILE)

seed.sql:
	$(GO) run ./cmd/gen -output=$(CURDIR)/seed.sql

music-api: coverage
	$(GO) build -o $@ ./cmd/api

music-gen: coverage
	$(GO) build -o $@ ./cmd/gen