package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	URL := "https://codember.dev/data/message_02.txt"

	var VALUE int = 0

	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var builder strings.Builder
	_, err = io.Copy(&builder, resp.Body)
	if err != nil {
		panic(err)
	}

	content := builder.String()

	for i := 0; i < len(content); i++ {
		item := string(content[i])

		if item == "#" {
			VALUE += 1
		} else if item == "@" {
			VALUE -= 1
		} else if item == "*" {
			VALUE *= VALUE
		} else if item == "&" {
			fmt.Print(VALUE)
		} else {
		}
	}
}
