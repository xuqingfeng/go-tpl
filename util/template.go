package util

import (
	"bytes"
	"text/template"
)

var (
	gitignore = `.DS_Store
.idea
out
{{ .Repo }}
`
	travis = `language: go
go:
  - 1.8
  - tip
os:
  - linux
install:
  - go get -v ./...
script:
  - make test
`
	makefile = `deps:
	go get -v -d ./...

test:
	go test -v $$(go list ./... | grep -v /vendor/)

fmt:
	go fmt ./...

build: fmt
	go build

run: build
	./{{ .Repo }}
`
	readme = `# {{ .Repo }}
[![travis](https://img.shields.io/travis/{{ .Author }}/{{ .Repo }}/master.svg?style=flat-square)](https://travis-ci.org/{{ .Author }}/{{ .Repo }})
[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/{{ .Author }}/{{ .Repo }})

### Install

` + "```bash" +
		`
go get github.com/{{ .Author }}/{{ .Repo }}
` + "```" +
		`
### Usage

### Test

` + "```bash" +
		`
make test
` + "```"

	main = `package main

import "fmt"

func main(){

    fmt.Print("Hello, Golang!")
}
`

	fs = map[string]string{
		".gitignore":  gitignore,
		".travis.yml": travis,
		"Makefile":    makefile,
		"README.md":   readme,
		"main.go":     main,
	}
)

// Data get input from cli
type Data struct {
	Author string
	Repo   string
}

// Feed convert template to content
func Feed(d Data) (map[string]string, error) {

	for name, content := range fs {
		parsed, err := parse(d, name, content)
		if err != nil {
			return nil, err
		}
		fs[name] = parsed
	}

	return fs, nil
}

func parse(d Data, name, content string) (string, error) {

	tpl, err := template.New(name).Parse(content)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, d)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
