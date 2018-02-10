BIN_NAME=conbot
BUILD_VERSION=latest

ci: clean devdeps dep lint build test

clean:
	go clean
	rm -f $(BIN_NAME)

devdeps:
	go get github.com/golang/dep/cmd/dep
	go get github.com/golang/lint/golint

dep:
	dep ensure

lint:
	go vet ./...
	golint -set_exit_status $$(go list ./...)

build:
	go build

test:
	go test ./...

fmt:
	gofmt -s -w .

run:
	./run.sh
