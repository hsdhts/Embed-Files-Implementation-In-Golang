package golang_embed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestName(t *testing.T) {
	fmt.Println(version)
}

//go:embed chisato.jpg
var logo []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("chisato_new.jpg", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
