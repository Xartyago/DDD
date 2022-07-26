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
	var ts []domain.Transaction
	if err := r.db.Read(&ts); err != nil {
		return nil, errors.New(web.ReadFile)
	}
	return ts, nil
}

// Create a new Transaction
func (r *repository) Store(lastId int, transactionCode, currency, emisor, receiver, transactionDate string, amount float64) (domain.Transaction, error) {
	var ts []domain.Transaction
	// Get the ts from the .json
	if err := r.db.Read(&ts); err != nil {
		return domain.Transaction{}, errors.New(web.ReadFile)
	}
	for _, tction := range ts {
		if transactionCode == tction.TransactionCode {
			return domain.Transaction{}, fmt.Errorf("the transaction already exists")
		}
	}
	// Create the new tction
	newts := domain.Transaction{
		Id:              lastId,
		TransactionCode: transactionCode,
		Currency:        currency,
		Emisor:          emisor,
		Receiver:        receiver,
		TransactionDate: transactionDate,
		Amount:          amount,
	}
	// Push into the .json
	ts = append(ts, newts)
	r.db.Write(ts)
	return newts, nil
}

// Update whole Transaction
func (r *repository) Update(idToUpdate int, transactionCode, currency, emisor, receiver, transactionDate string, amount float64) (domain.Transaction, error) {
	// Get the transaction from the .json
	var ts []domain.Transaction
	if err := r.db.Read(&ts); err != nil {
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
	// Update transaction into the []ts
	exist := false
	for i := range ts {
		if ts[i].Id == idToUpdate {
			transactionUpdated.Id = idToUpdate
			ts[i] = transactionUpdated
			exist = true
		}
	}
	if !exist {
		return domain.Transaction{}, fmt.Errorf("the transaction with the id %d has not been found", idToUpdate)
	}
	// Push into the .json
	r.db.Write(ts)
	return transactionUpdated, nil
}

// Patch Transaction Code
func (r *repository) PatchCode(idToPatch int, transactionCode string) (domain.Transaction, error) {
	// Get the transaction from the .json
	var ts []domain.Transaction
	if err := r.db.Read(&ts); err != nil {
		return domain.Transaction{}, errors.New(web.ReadFile)
	}
	// Update the transaction
	var transactionPatched domain.Transaction
	exist := false
	for i := range ts {
		if ts[i].Id == idToPatch {
			ts[i].TransactionCode = transactionCode
			transactionPatched = ts[i]
			exist = true
		}
	}
	if !exist {
		return domain.Transaction{}, fmt.Errorf("the transaction with the id %d has not been found", idToPatch)
	}
	// Push into the .json
	r.db.Write(ts)
	return transactionPatched, nil
}

// Patch Amount
func (r *repository) PatchAmount(idToPatch int, amount float64) (domain.Transaction, error) {
	// Get the transaction from the .json
	var ts []domain.Transaction
	if err := r.db.Read(&ts); err != nil {
		return domain.Transaction{}, errors.New(web.ReadFile)
	}
	// Update the transaction
	var transactionPatched domain.Transaction
	exist := false
	for i := range ts {
		if ts[i].Id == idToPatch {
			ts[i].Amount = amount
			transactionPatched = ts[i]
			exist = true
		}
	}
	if !exist {
		return domain.Transaction{}, fmt.Errorf("the transaction with the id %d has not been found", idToPatch)
	}
	// Push into the .json
	r.db.Write(ts)
	return transactionPatched, nil
}

// Delete the transaction with the id specificed
func (r *repository) Delete(idToDelete int) (domain.Transaction, error) {
	// Get the transaction from the .json
	var ts []domain.Transaction
	if err := r.db.Read(&ts); err != nil {
		return domain.Transaction{}, errors.New(web.ReadFile)
	}
	// Delete transaction
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
	// Push into the .json
	r.db.Write(ts)
	return transactionDeleted, nil
}

// Get the last id
func (r *repository) LastId() (int, error) {
	var ts []domain.Transaction
	if err := r.db.Read(&ts); err != nil {
		return 0, errors.New(web.ReadFile)
	}
	lastId := ts[len(ts)-1].Id + 1
	return lastId, nil
}
