package md2html

import (
	"bytes"
	"github.com/Yoak3n/gulu/logger"
	"github.com/Yoak3n/gulu/util"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"

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
	source, _ := os.ReadFile(name)
	return Md2BlockDOM(engine, source)
}

func DirMd2HTML(engine *goldmark.Markdown, rootDir string, outDir string) ([]string, error) {
	files, err := extractMdFiles(rootDir)
	if err != nil {
		return nil, err
	}
	wg := &sync.WaitGroup{}
	wg.Add(len(files))
	for _, file := range files {
		go func() {
			html, e := MdFile2HTML(engine, file)
			if e != nil {
				logger.Logger.Errorln(e)
			}
			targetPath, _ := strings.CutSuffix(file, ".md")
			targetPath = targetPath + ".html"
			targetPath = filepath.Join(outDir, targetPath)
			targeDir := filepath.Dir(targetPath)
			util.CreateDirNotExists(targeDir)
			logger.Logger.Println(targetPath)
			fp, e := os.Create(targetPath)
			if e != nil {
				logger.Logger.Errorln(e)
			}
			_, e = fp.WriteString(html)
			if e != nil {
				return
			}
			defer wg.Done()
			defer fp.Close()
		}()

	}
	wg.Wait()
	return files, nil
}

func extractMdFiles(dir string) ([]string, error) {
	files := make([]string, 0)
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".md" {
			replacePath := strings.ReplaceAll(path, " ", "-")
			if err != nil {
				return err
			}
			files = append(files, replacePath)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files, nil
}
