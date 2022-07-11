package transactions

import (
	"fmt"

	"github.com/Xartyago/DDD/internal/domain"
	"github.com/Xartyago/DDD/pkg/store"
)

type Repository interface {
	GetAll() ([]domain.Transaction, error)
	Store(transactionCode, currency, emisor, receiver, transactionDate string, amount float64) (domain.Transaction, error)
	Update(idToUpdate int, transactionCode, currency, emisor, receiver, transactionDate string, amount float64) (domain.Transaction, error)
	PatchCode(idToPatch int, transactionCode string) (domain.Transaction, error)
	PatchAmount(idToPatch int, amount float64) (domain.Transaction, error)
	Delete(idToDelete int) (domain.Transaction, error)
}

type repository struct {
	db store.Store
}

var ts []domain.Transaction

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

// Get all Transactions
func (r *repository) GetAll() ([]domain.Transaction, error) {
	var tctions []domain.Transaction
	if err := r.db.Read(&tctions); err != nil {
		return nil, fmt.Errorf("cant read the databse")
	}
	return tctions, nil
}

// Create a new Transaction
func (r *repository) Store(transactionCode, currency, emisor, receiver, transactionDate string, amount float64) (domain.Transaction, error) {
	var tctions []domain.Transaction
	// Get the tctions from the .json
	if err := r.db.Read(&tctions); err != nil {
		return domain.Transaction{}, fmt.Errorf("cant read the databse")
	}
	for _, tction := range tctions {
		if transactionCode == tction.TransactionCode {
			return domain.Transaction{}, fmt.Errorf("the transaction already exists")
		}
	}
	lastId := tctions[len(tctions)-1].Id + 1
	// Create the new tction
	newTs := domain.Transaction{
		Id:              lastId,
		TransactionCode: transactionCode,
		Currency:        currency,
		Emisor:          emisor,
		Receiver:        receiver,
		TransactionDate: transactionDate,
		Amount:          amount,
	}
	// Push into the .json
	tctions = append(tctions, newTs)
	r.db.Write(tctions)
	return newTs, nil
}

// Update whole Transaction
func (r *repository) Update(idToUpdate int, transactionCode, currency, emisor, receiver, transactionDate string, amount float64) (domain.Transaction, error) {
	tsUpdated := domain.Transaction{
		TransactionCode: transactionCode,
		Currency:        currency,
		Emisor:          emisor,
		Receiver:        receiver,
		TransactionDate: transactionDate,
		Amount:          amount,
	}
	exist := false
	for i := range ts {
		if ts[i].Id == idToUpdate {
			tsUpdated.Id = idToUpdate
			ts[i] = tsUpdated
			exist = true
		}
	}
	if !exist {
		return domain.Transaction{}, fmt.Errorf("the transaction with the id %d has not been found", idToUpdate)
	}
	return tsUpdated, nil
}

// Patch Transaction Code
func (r *repository) PatchCode(idToPatch int, transactionCode string) (domain.Transaction, error) {
	exist := false
	for i := range ts {
		if ts[i].Id == idToPatch {
			ts[i].TransactionCode = transactionCode
			exist = true
		}
	}
	if !exist {
		return domain.Transaction{}, fmt.Errorf("the transaction with the id %d has not been found", idToPatch)
	}
	return ts[idToPatch], nil
}

// Patch Amount
func (r *repository) PatchAmount(idToPatch int, amount float64) (domain.Transaction, error) {
	exist := false
	for i := range ts {
		if ts[i].Id == idToPatch {
			ts[i].Amount = amount
			exist = true
		}
	}
	if !exist {
		return domain.Transaction{}, fmt.Errorf("the transaction with the id %d has not been found", idToPatch)
	}
	return ts[idToPatch], nil
}

// Delete the transaction with the id specified
func (r *repository) Delete(idToDelete int) (domain.Transaction, error) {
	var transactionDeleted domain.Transaction
	exist := false
	for i := range ts {
		if ts[i].Id == idToDelete {
			transactionDeleted = ts[i]
			ts = append(ts[:idToDelete], ts[idToDelete+1:]...)
			exist = true
		}
	}
	if !exist {
		return domain.Transaction{}, fmt.Errorf("the id %d doesn't exist", idToDelete)
	}
	return transactionDeleted, nil
}
