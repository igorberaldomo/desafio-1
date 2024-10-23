package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	client := http.Client{Timeout:time.Microsecond*200 }
	resp, err := client.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err :=io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}