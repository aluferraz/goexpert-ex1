package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"projeto_1/models"
	"time"
)

func main() {
	GetDollarRate()
}

func GetDollarRate() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Println("Erro ao criar request")
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Erro ao executar request")
		panic(err)
	}
	defer res.Body.Close()
	var bid models.BidResponse
	body, err := io.ReadAll(res.Body) // response body is []byte

	err = json.Unmarshal(body, &bid)
	if err != nil {
		log.Println("Erro ao interpretar body")
	}
	StoreBID(&bid)
	log.Printf("Dólar %v\n", bid.Bid)

}

func StoreBID(bid *models.BidResponse) {
	f, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}
	_, err = fmt.Fprintf(f, "Dólar %v\n", bid.Bid)
	if err != nil {
		log.Println("Erro ao gravar arquivo")
		panic(err)
	}
	defer f.Close()

}
