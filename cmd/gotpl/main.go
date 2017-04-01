package main

import (
	"flag"
	"log"

	"github.com/xuqingfeng/gotpl/util"
)

func main() {

	data := util.Data{}
	flag.StringVar(&data.Author, "author", "", "author of the porject")
	flag.StringVar(&data.Repo, "repo", "", "repository name")

	flag.Parse()

	if data.Author == "" || data.Repo == "" {
		log.Fatalln("E! author or repo can't be empty")
	}

	err := util.CreateDirectory(data)
	if err != nil {
		log.Fatalf("E! %v", err)
	}
	newFs, err := util.Feed(data)
	if err != nil {
		log.Fatalf("E! %v", err)
	}
	err = util.CreateFile(data, newFs)
	if err != nil {
		log.Fatalf("E! %v", err)
	}
}
