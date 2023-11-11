package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	URL := "https://codember.dev/data/message_01.txt"

	resp, err := http.Get(URL)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var builder strings.Builder
	_, err = io.Copy(&builder, resp.Body)
	if err != nil {
		fmt.Println("Error al leer el cuerpo de la respuesta:", err)
		return
	}

	data := builder.String()
	items := strings.Split(data, " ")
	maps := make(map[string]int)
	slice := []string{}
	insert := func(key string, value int) {
		if _, exist := maps[key]; !exist {
			maps[key] = value
			slice = append(slice, key)
		} else {
			if maps[key] != value {
				maps[key] = value
			}
		}
	}

	for _, item := range items {
		// serialization
		item = strings.ToLower(strings.TrimSpace(item))
		if _, err := maps[item]; !err {
			insert(item, 1)
		} else {
			insert(item, maps[item]+1)
		}
	}

	for _, key := range slice {
		fmt.Printf("%s%d", key, maps[key])
	}
}
