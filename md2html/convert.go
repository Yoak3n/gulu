package md2html

import (
	"bytes"
	"fmt"
	"os"
	"path"

	"github.com/yuin/goldmark"
)

func Md2BlockDOM(engine *goldmark.Markdown, md []byte) (string, error) {
	var buf bytes.Buffer

	if err := goldmark.Convert(md, &buf); err != nil {
		return "", err
	} else {
		return buf.String(), nil
	}
}

func MdFile2HTML(engine *goldmark.Markdown, name string) (string, error) {
	title := path.Base(name)
	fmt.Println("title:", title)
	source, _ := os.ReadFile("./bing-qqbot.md")
	return Md2BlockDOM(engine, source)
}
