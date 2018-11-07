// +build linux

package main

import (
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
