// +build windows

package main

import (
	"path/filepath"
)

func pathify(path string) string {
	return filepath.Clean(path[1:])
}
