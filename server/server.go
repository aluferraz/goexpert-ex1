package main

import (
	"context"
	"encoding/json"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"projeto_1/models"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/cotacao", HandleGetRate)
	http.ListenAndServe(":8080", mux)
}

func HandleGetRate(w http.ResponseWriter, r *http.Request) {
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")
	rate := GetDollarRate()
	storeDollarRate(&rate.Rate)
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&models.BidResponse{Bid: rate.Bid})
	if err != nil {
		println("Erro ao escrever resposta")
		return
	}
}

func storeDollarRate(rate *models.Rate) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	db, err := gorm.Open(sqlite.Open("rates.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	tx := db.WithContext(ctx)
	// Migrate the schema (for simplicity, let's just call this here (since we don't have a CI/CD yet))
	tx.AutoMigrate(&models.Rate{})
	// Create
	err = tx.Create(rate).Error
	if err != nil {
		log.Println("Erro ao salvar")
		panic(err)
	}
}

func GetDollarRate() models.DollarRate {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		log.Println("Erro ao criar request")
		panic(err)

	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Erro ao consultar https://economia.awesomeapi.com.br/json/last/USD-BRL")
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Erro ao ler body")
		panic(err)
	}
	var dolarRate models.DollarRate
	err = json.Unmarshal(body, &dolarRate)
	if err != nil {
		log.Println("Erro ao ler body")
		panic(err)
	}
	//io.Copy(os.Stdout, res.Body)
	return dolarRate
}
