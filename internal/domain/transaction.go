package domain

type Transaction struct {
	Id              int     `json:"id"`
	TransactionCode string  `json:"transactioncode"`
	Currency        string  `json:"currency"`
	Emisor          string  `json:"emisor"`
	Receiver        string  `json:"receiver"`
	TransactionDate string  `json:"transactiondate"`
	Amount          float64 `json:"amount"`
}
