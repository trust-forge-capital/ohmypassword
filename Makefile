VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILDTIME := $(shell date -u '+%Y-%m-%d %H:%M:%S')
GITCOMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

LDFLAGS := -X github.com/trust-forge-capital/ohmypassword/cmd/cli.version=$(VERSION) \
           -X github.com/trust-forge-capital/ohmypassword/cmd/cli.buildTime=$(BUILDTIME) \
           -X github.com/trust-forge-capital/ohmypassword/cmd/cli.gitCommit=$(GITCOMMIT)

run:
	go run ./cmd/ohmypassword

build:
	go build -ldflags="$(LDFLAGS)" -o bin/ohmypassword ./cmd/ohmypassword

build-all:
	GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o bin/ohmypassword-linux-amd64 ./cmd/ohmypassword
	GOOS=linux GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o bin/ohmypassword-linux-arm64 ./cmd/ohmypassword
	GOOS=darwin GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o bin/ohmypassword-darwin-amd64 ./cmd/ohmypassword
	GOOS=darwin GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o bin/ohmypassword-darwin-arm64 ./cmd/ohmypassword
	GOOS=windows GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o bin/ohmypassword-windows-amd64.exe ./cmd/ohmypassword

test:
	go test -v -race -coverprofile=coverage.txt ./...

lint:
	golangci-lint run ./...

fmt:
	go fmt ./...
	gofmt -s -w .

clean:
	rm -rf bin/
	rm -f coverage.txt

install:
	go install ./cmd/ohmypassword

.PHONY: run build build-all test lint fmt clean install