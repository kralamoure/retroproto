package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var rx = regexp.MustCompile(`^// generated \(unrevised\)`)

func main() {
	dirnames := []string{"msgsvr", "msgcli"}

	for _, dirname := range dirnames {
		err := processDir(dirname)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
}

func processDir(dirname string) error {
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		return err
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}
		ext := filepath.Ext(fileInfo.Name())
		ext = strings.TrimPrefix(ext, ".")
		if ext != "go" {
			continue
		}

		err = processFile(filepath.Join(dirname, fileInfo.Name()))
		if err != nil {
			return err
		}
	}

	return nil
}

func processFile(filename string) error {
	p, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	match := rx.Match(p)
	if !match {
		return nil
	}

	err = os.Remove(filename)
	if err != nil {
		return err
	}

	return nil
}
