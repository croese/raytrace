all: test build

build:
	go build -o cannon ./cmd/cannon
	go build -o rt ./cmd/rt

test:
	go test -v ./...