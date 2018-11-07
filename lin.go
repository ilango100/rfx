// +build linux

package main

import (
	"io/ioutil"
	"net/url"
	"path/filepath"
)

func pathify(path string) string {
	uri, err := url.ParseRequestURI(path)
	if err != nil {
		return ""
	}
	return filepath.Clean(uri.Path)
}

var listDir = ioutil.ReadDir
