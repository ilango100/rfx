package main

import (
	"crypto"
	_ "crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
)

func init() {
	if err := loadSettings(); err != nil {
		firstStart()
	}
}

func firstStart() {
	fmt.Println(welcomestr)
	fmt.Print("Enable authentication? (y/n) ")
	var auth string
	fmt.Scan(&auth)
	auth = strings.ToLower(auth)
	if auth == "y" || auth == "yes" {
		set.Auth = true
	}
	if set.Auth {
		fmt.Print("Enter password: ")
		var pass string
		fmt.Scan(&pass)
		if crypto.MD5.Available() {
			md5 := crypto.MD5.New()
			md5.Write([]byte(pass))
			set.Pass = hex.EncodeToString(md5.Sum([]byte{}))
		} else {
			log.Fatalln("Cannot find MD5")
		}
	}

	fmt.Print("Port number to listen (80): ")
	fmt.Scan(&set.Port)
	if set.Port == 0 {
		set.Port = 80
	}

	fmt.Printf("%#v\n", set)
	if err := saveSettings(); err != nil {
		log.Fatalln(err)
	}
}
