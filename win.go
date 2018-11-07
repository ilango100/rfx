// +build windows

package main

import (
	"io/ioutil"
	"os"
)

func listDir(dir string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dir[1:])
}
