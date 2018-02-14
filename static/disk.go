// +build devel

package static

import (
	"io/ioutil"
	"path"
)

func File(file string) []byte {
	r, err := ioutil.ReadFile(path.Join("static", file))
	if err != nil {
		panic(err)
	}
	return r
}
