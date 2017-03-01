
setup:
	go get -u github.com/jteeuwen/go-bindata/...
.PHONY: setup

generate:
	go generate ./...
.PHONY: generate
