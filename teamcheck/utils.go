package main

import (
	"net/url"
	"os"
	"log"
)

func CheckServer(server string) bool {

	_, err := url.ParseRequestURI(server)
	if err != nil {
		return false
	} else {
		return true
	}
}

func GetExecutable() string {
	ex, e := os.Executable()
	if e != nil {
		log.Fatal("Couldn't get executable process")
	}
	return ex
}

