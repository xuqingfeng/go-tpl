package main

import (
	"flag"
	"log"

	"github.com/xuqingfeng/go-tpl/util"
)

func main() {

	data := util.Data{}
	flag.StringVar(&data.Author, "author", "author", "author of the porject")
	flag.StringVar(&data.Repo, "repo", "repo", "repository name")

	flag.Parse()

	err := util.CreateDirectory(data)
	if err != nil {
		log.Fatal(err)
	}
	newFs, err := util.Feed(data)
	if err != nil {
		log.Fatal(err)
	}
	err = util.CreateFile(data, newFs)
	if err != nil {
		log.Fatal(err)
	}
}
