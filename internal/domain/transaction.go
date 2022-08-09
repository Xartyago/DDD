package domain

type Transaction struct {
	Id              int     `json:"id"`
	TransactionCode string  `json:"transaction_code"`
	Currency        string  `json:"currency"`
	Emisor          string  `json:"emisor"`
	Receiver        string  `json:"receiver"`
	TransactionDate string  `json:"transaction_date"`
	Amount          float64 `json:"amount"`
}
