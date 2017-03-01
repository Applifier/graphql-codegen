
setup:
	go get -u github.com/jteeuwen/go-bindata/...
.PHONY: setup

generate:
	go generate ./...
.PHONY: generate

test: generate
	go test ./...
.PHONY: test

record: generate
	RECORD_FIXTURES=yes go test ./...
.PHONY: record
