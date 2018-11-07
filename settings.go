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
}

func getHome() string {
	usr, err := user.Current()
	if err != nil {
		return ""
	}
	return usr.HomeDir
}

func saveSettings() error {
	savepath := path.Join(getHome(), "settings.json")
	fl, err := os.Create(savepath)
	if err != nil {
		return err
	}
	jb, err := json.Marshal(set)
	if err != nil {
		return err
	}
	_, err = fl.Write(jb)
	return err
}
