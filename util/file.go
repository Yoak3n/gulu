package util

import (
	"os"
	"strings"
)

func CreateDirNotExists(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		e := os.MkdirAll(dir, os.ModePerm)
		if e != nil {
			println("Error creating directory: " + e.Error())
			return
		}
	}
}
func CreateFileNotExists(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		f, e := os.Create(filePath)
		if e != nil {
			return e
		}
		err = f.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func GetFileDir(filePath string) string {
	if strings.HasSuffix(filePath, "/") {
		filePath, _ = strings.CutSuffix(filePath, "/")
	}
	return filePath[strings.LastIndex(filePath, "/")+1:]
}
