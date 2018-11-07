// +build windows

package main

import (
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
)

var drives []os.FileInfo

func pathify(path string) string {
	uri, err := url.ParseRequestURI(path)
	if err != nil {
		return ""
	}
	return filepath.Clean(uri.Path[1:])
}

func listDir(path string) ([]os.FileInfo, error) {
	if path == "" {
		return drives, nil
	}
	return ioutil.ReadDir(path)
}
