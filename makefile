
all: build

check:
	type go

build: check test
	go build -v -o ./bin/zsh-trophy

test: check
	go test -v ./...
