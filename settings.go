package main

import (
	"encoding/json"
	"os"
	"os/user"
	"path"
)

var set Settings

//Settings allows changing preferences for the app
type Settings struct {
	Auth    bool   `json:"auth"`
	Pass    string `json:"pass"`
	Backend bool   `json:"back"`
	Port    int    `json:"port"`
}

func getHome() string {
	usr, err := user.Current()
	if err != nil {
		return ""
	}
	return usr.HomeDir
}

func saveSettings() error {
	savepath := path.Join(getHome(), setfilename)
	fl, err := os.Create(savepath)
	if err != nil {
		return err
	}
	defer fl.Close()
	jsn := json.NewEncoder(fl)
	jsn.SetIndent("", " ")
	return jsn.Encode(set)
}

func loadSettings() error {
	savepath := path.Join(getHome(), setfilename)
	fl, err := os.Open(savepath)
	if err != nil {
		return err
	}
	defer fl.Close()
	jsn := json.NewDecoder(fl)
	return jsn.Decode(&set)
}
