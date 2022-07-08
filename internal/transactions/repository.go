package transactions

import (
	"github.com/Xartyago/DDD/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.Transaction, error)
	Store(id int, transactionCode, currency, emisor, receiver, transactionDate string, amount float64) (domain.Transaction, error)
}

type repository struct{}

var ts []domain.Transaction

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]domain.Transaction, error) {
	// trctionsJson, readJsonErr := os.Open("./transactions.json")
	// fmt.Println(trctionsJson)
	// if readJsonErr != nil {
	// 	return nil, readJsonErr
	// }
	// var parseTsBytes []domain.Transaction
	// if err := json.Unmarshal(trctionsJson, &parseTsBytes); err != nil {
	// 	return nil, err
	// }
	// ts = parseTsBytes
	return ts, nil
}

func (r *repository) Store(id int, transactionCode, currency, emisor, receiver, transactionDate string, amount float64) (domain.Transaction, error) {
	newTs := domain.Transaction{
		Id:              id,
		TransactionCode: transactionCode,
		Currency:        currency,
		Emisor:          emisor,
		Receiver:        receiver,
		TransactionDate: transactionDate,
		Amount:          amount,
	}
	ts = append(ts, newTs)
	return newTs, nil
}
