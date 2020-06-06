package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/kralamoure/d1proto"
)

func main() {
	err := os.MkdirAll(filepath.Join("out", "msgcli"), 0775)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(filepath.Join("out", "msgsvr"), 0775)
	if err != nil {
		panic(err)
	}

	msgSvrTmpl, err := template.ParseFiles("assets/template.msgsvr.txt")
	if err != nil {
		panic(err)
	}

	for _, v := range d1proto.MsgSvrIds {
		name, _ := d1proto.MsgSvrNameByID(v)
		buf := &bytes.Buffer{}
		err := msgSvrTmpl.Execute(buf, name)
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(
			filepath.Join("out", "msgsvr",
				fmt.Sprintf("%s.go", strings.ToLower(name)),
			), buf.Bytes(), 0664)
		if err != nil {
			panic(err)
		}
	}

	msgCliTmpl, err := template.ParseFiles("assets/template.msgcli.txt")
	if err != nil {
		panic(err)
	}

	for _, v := range d1proto.MsgCliIds {
		name, _ := d1proto.MsgCliNameByID(v)
		buf := &bytes.Buffer{}
		err := msgCliTmpl.Execute(buf, name)
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(
			filepath.Join("out", "msgcli",
				fmt.Sprintf("%s.go", strings.ToLower(name)),
			), buf.Bytes(), 0664)
		if err != nil {
			panic(err)
		}
	}
}
