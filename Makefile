GO_PACKAGES=$(shell go list ./... | grep -v vendor)

.PHONY: all \
        build \
        fmt \
        generate-golang-schema \
        lint \
        tests \
        vet

all: generate-golang-schema fmt lint build tests

build:
	go build ./...

generate-golang-schema:
	rm -f pkg/claim/schema.go
	go run cmd/generate/generate.go

fmt:
	go fmt ./...

lint:
	golangci-lint run

tests:
	go test -coverprofile=cover.out ./...

validate-example:
	jsonschema -i schemas/claim.example.json schemas/claim.schema.json

vet:
	go vet ${GO_PACKAGES}
