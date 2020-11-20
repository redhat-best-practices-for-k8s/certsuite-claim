.PHONY: all \
        build \
        fmt \
        generate-golang-schema \
        lint \
        tests

# Export GO111MODULE=on to enable project to be built from within GOPATH/src
export GO111MODULE=on

all: generate-golang-schema fmt lint build tests

build:
	go build ./...

generate-golang-schema:
	rm -f pkg/claim/schema.go
	go run cmd/generate/generate.go

fmt:
	go fmt ./...

lint:
	golint `go list ./... | grep -v vendor`
	golangci-lint run

tests:
	go test -coverprofile=cover.out ./...
