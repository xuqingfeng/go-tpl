deps:
	go get -v -d ./...

test:
	go vet ./... && go test -v $$(go list ./... | grep -v /vendor/)

fmt:
	go fmt ./...

build: fmt
	go build -o ./out/gotpl ./cmd/gotpl

run: build
	./out/gotpl
