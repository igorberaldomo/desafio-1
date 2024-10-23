package main

import (
	"io"
	"net/http"
	"time"
	"context"
)

func main() {
	// server disponibiliza as informções da API
	http.HandleFunc("/cotacao", handleCotacao)
	http.ListenAndServe(":8080", nil)
}

func handleCotacao(w http.ResponseWriter, r *http.Request) {
	// faz a requisição para a API
	ctxreq, cancel := context.WithTimeout(r.Context(), time.Millisecond*200)
	defer cancel()
	// server consome a API
	req, err := http.NewRequestWithContext(ctxreq, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	res , err := http.DefaultClient.Do(req)
	if err != nil {
	}
	defer res.Body.Close()
	body, err :=io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}