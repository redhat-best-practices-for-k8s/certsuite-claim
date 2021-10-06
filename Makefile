.PHONY: all \
        build \
        fmt \
        generate-golang-schema \
        lint \
        tests

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

validate-example:
	jsonschema -i schemas/claim.example.json schemas/claim.schema.json
