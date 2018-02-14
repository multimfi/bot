// +build ignore

package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
	"unsafe"
)

const fileTemplate = `// +build !devel

package static

import "log"
{{ range . }}
var {{ .Var }} = {{ .Bytes | printf "%#v" }}
{{ end }}
func File(file string) []byte {
	switch file {
{{ range . }}
	case {{ . | printf "%q" }}:
		return {{ .Var }}
{{ end }}
	default:
		log.Fatalln(file, "not found")
	}

	return *(*[]byte)(nil)
}
`

type file string

type files []file

func (f file) Var() string {
	return strings.Replace(string(f), ".", "_", -1)
}

func (f file) Bytes() []byte {
	r, err := ioutil.ReadFile(string(f))
	if err != nil {
		log.Fatal(err)
	}
	return r
}

func main() {
	t := template.Must(template.New("").Parse(fileTemplate))
	f := *(*files)(unsafe.Pointer(&os.Args))
	if err := t.Execute(os.Stdout, f[1:]); err != nil {
		log.Fatal(err)
	}
}
