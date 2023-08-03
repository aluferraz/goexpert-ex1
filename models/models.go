package models

//Usando models compartilhados entre cliente e servidor para praticar go modules
import "time"

type Rate struct {
	ID         uint
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type BidResponse struct {
	Bid string `json:"bid"`
}

type DollarRate struct {
	Rate `json:"USDBRL"`
}
