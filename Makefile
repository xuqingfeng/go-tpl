deps:
	go get -v -d ./...

test:
	go test -v $$(go list ./... | grep -v /vendor/)

fmt:
	go fmt ./...

build: fmt
	go build -o ./out/go-tpl ./cmd

run: build
	./out/go-tpl
