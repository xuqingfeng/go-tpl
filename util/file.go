package util

import (
	"io/ioutil"
	"os"
)

var (
	repo = `{{ .Repo }}`
	out  = `{{ .Repo }}/out`
	dir  = map[string]string{
		"repo": repo,
		"out":  out,
	}
)

const (
	dirMode  = 0755
	fileMode = 0644
)

func CreateDirectory(d Data) error {

	for k, v := range dir {
		parsed, err := parse(d, k, v)
		if err != nil {
			return err
		}
		err = os.Mkdir(parsed, dirMode)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateFile(d Data, fs map[string]string) error {

	for name, content := range fs {
		err := ioutil.WriteFile(d.Repo+"/"+name, []byte(content), fileMode)
		if err != nil {
			return err
		}
	}

	return nil
}
