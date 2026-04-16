run:
	go run ./cmd/cli

build:
	go build -o bin/ohmypassword ./cmd/cli

build-all:
	GOOS=linux GOARCH=amd64 go build -o bin/ohmypassword-linux-amd64 ./cmd/cli
	GOOS=linux GOARCH=arm64 go build -o bin/ohmypassword-linux-arm64 ./cmd/cli
	GOOS=darwin GOARCH=amd64 go build -o bin/ohmypassword-darwin-amd64 ./cmd/cli
	GOOS=darwin GOARCH=arm64 go build -o bin/ohmypassword-darwin-arm64 ./cmd/cli
	GOOS=windows GOARCH=amd64 go build -o bin/ohmypassword-windows-amd64.exe ./cmd/cli

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
	go install ./cmd/cli

.PHONY: run build build-all test lint fmt clean install