
setup:
	go get -u github.com/jteeuwen/go-bindata/...
.PHONY: setup

generate:
	go generate ./...
.PHONY: generate

test:
	go test ./...
.PHONY: test

record:
	RECORD_FIXTURES=yes go test ./...
.PHONY: record
