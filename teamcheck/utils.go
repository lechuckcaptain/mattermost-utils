package main

import (
	"net/url"
)

func CheckServer(server string) bool {

	_, err := url.ParseRequestURI(server)
	if err != nil {
		return false
	} else {
		return true
	}
}

