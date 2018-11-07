// +build windows

package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

var drives []os.FileInfo

type drive struct {
	name  string
	finfo os.FileInfo
}

func (d drive) Name() string {
	return d.name
}

func (d drive) Size() int64 {
	return d.finfo.Size()
}
func (d drive) Mode() os.FileMode {
	return d.finfo.Mode()
}
func (d drive) ModTime() time.Time {
	return d.finfo.ModTime()
}
func (d drive) IsDir() bool {
	return d.finfo.IsDir()
}
func (d drive) Sys() interface{} {
	return d.finfo.Sys()
}

func init() {
	for _, d := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		st, err := os.Stat(fmt.Sprintf("%c:", d))
		if err != nil {
			continue
		}
		drives = append(drives, drive{
			name:  fmt.Sprintf("/%c:", d),
			finfo: st,
		})
	}
	fmt.Println(drives)
	for _, d := range drives {
		fmt.Println(d.Name())
	}
}

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
