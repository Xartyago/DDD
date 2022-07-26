package transactions

import (
	"errors"
	"fmt"

	"github.com/Xartyago/DDD/internal/domain"
	"github.com/Xartyago/DDD/pkg/store"
	"github.com/Xartyago/DDD/pkg/web"
)

type Repository interface {
	GetAll() ([]domain.Transaction, error)
	Store(lastId int, transactionCode, currency, emisor, receiver, transactionDate string, amount float64) (domain.Transaction, error)
	Update(idToUpdate int, transactionCode, currency, emisor, receiver, transactionDate string, amount float64) (domain.Transaction, error)
	PatchCode(idToPatch int, transactionCode string) (domain.Transaction, error)
	PatchAmount(idToPatch int, amount float64) (domain.Transaction, error)
	Delete(idToDelete int) (domain.Transaction, error)
	LastId() (int, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

// Get all Transactions
func (r *repository) GetAll() ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	if err := r.db.Read(&transactions); err != nil {
		return nil, errors.New(web.ReadFile)
	}
	return transactions, nil
}

// Create a new Transaction
func (r *repository) Store(lastId int, transactionCode, currency, emisor, receiver, transactionDate string, amount float64) (domain.Transaction, error) {
	var transactions []domain.Transaction
	// Get the transactions from the .json
	if err := r.db.Read(&transactions); err != nil {
		return domain.Transaction{}, errors.New(web.ReadFile)
	}
	for _, tction := range transactions {
		if transactionCode == tction.TransactionCode {
			return domain.Transaction{}, fmt.Errorf("the transaction already existransactions")
		}
	}
	// Create the new tction
	newtransactions := domain.Transaction{
		Id:              lastId,
		TransactionCode: transactionCode,
		Currency:        currency,
		Emisor:          emisor,
		Receiver:        receiver,
		TransactionDate: transactionDate,
		Amount:          amount,
	}
	// Push into the .json
	transactions = append(transactions, newtransactions)
	r.db.Write(transactions)
	return newtransactions, nil
}

// Update whole Transaction
func (r *repository) Update(idToUpdate int, transactionCode, currency, emisor, receiver, transactionDate string, amount float64) (domain.Transaction, error) {
	// Get the transaction from the .json
	var transactions []domain.Transaction
	if err := r.db.Read(&transactions); err != nil {
		return domain.Transaction{}, errors.New(web.ReadFile)
	}
	// Create a transaction
	transactionUpdated := domain.Transaction{
		TransactionCode: transactionCode,
		Currency:        currency,
		Emisor:          emisor,
		Receiver:        receiver,
		TransactionDate: transactionDate,
		Amount:          amount,
	}
	// Update transaction into the []transactions
	exist := false
	for i := range transactions {
		if transactions[i].Id == idToUpdate {
			transactionUpdated.Id = idToUpdate
			transactions[i] = transactionUpdated
			exist = true
		}
	}
	if !exist {
		return domain.Transaction{}, fmt.Errorf("the transaction with the id %d has not been found", idToUpdate)
	}
	// Push into the .json
	r.db.Write(transactions)
	return transactionUpdated, nil
}

// Patch Transaction Code
func (r *repository) PatchCode(idToPatch int, transactionCode string) (domain.Transaction, error) {
	// Get the transaction from the .json
	var transactions []domain.Transaction
	if err := r.db.Read(&transactions); err != nil {
		return domain.Transaction{}, errors.New(web.ReadFile)
	}
	// Update the transaction
	exist := false
	for i := range transactions {
		if transactions[i].Id == idToPatch {
			transactions[i].TransactionCode = transactionCode
			exist = true
		}
	}
	if !exist {
		return domain.Transaction{}, fmt.Errorf("the transaction with the id %d has not been found", idToPatch)
	}
	// Push into the .json
	r.db.Write(transactions)
	return transactions[idToPatch], nil
}

// Patch Amount
func (r *repository) PatchAmount(idToPatch int, amount float64) (domain.Transaction, error) {
	// Get the transaction from the .json
	var transactions []domain.Transaction
	if err := r.db.Read(&transactions); err != nil {
		return domain.Transaction{}, errors.New(web.ReadFile)
	}
	// Update the transaction
	exist := false
	for i := range transactions {
		if transactions[i].Id == idToPatch {
			transactions[i].Amount = amount
			exist = true
		}
	}
	if !exist {
		return domain.Transaction{}, fmt.Errorf("the transaction with the id %d has not been found", idToPatch)
	}
	// Push into the .json
	r.db.Write(transactions)
	return transactions[idToPatch], nil
}

// Delete the transaction with the id specificed
func (r *repository) Delete(idToDelete int) (domain.Transaction, error) {
	// Get the transaction from the .json
	var transactions []domain.Transaction
	if err := r.db.Read(&transactions); err != nil {
		return domain.Transaction{}, errors.New(web.ReadFile)
	}
	// Delete transaction
	var transactionDeleted domain.Transaction
	exist := false
	for i := range transactions {
		if transactions[i].Id == idToDelete {
			transactionDeleted = transactions[i]
			transactions = append(transactions[:idToDelete], transactions[idToDelete+1:]...)
			exist = true
		}
	}
	if !exist {
		return domain.Transaction{}, fmt.Errorf("the id %d doesn't exist", idToDelete)
	}
	// Push into the .json
	r.db.Write(transactions)
	return transactionDeleted, nil
}

// Get the last id
func (r *repository) LastId() (int, error) {
	var transactions []domain.Transaction
	if err := r.db.Read(&transactions); err != nil {
		return 0, errors.New(web.ReadFile)
	}
	lastId := transactions[len(transactions)-1].Id + 1
	return lastId, nil
}
