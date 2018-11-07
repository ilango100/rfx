package main

import "fmt"

func init() {
	if err := loadSettings(); err != nil {
		firstStart()
	}
}

func firstStart() {
	fmt.Println(welcomestr)
}
