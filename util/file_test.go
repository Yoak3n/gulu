package util

import "testing"

func TestCreateFileNotExists(t *testing.T) {
	err := CreateFileNotExists("data/log/test.log", []byte("hello world"), []byte("hello world"))
	if err != nil {
		t.Error(err)
	}
}
