package main

import (
	"os/user"
)

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
