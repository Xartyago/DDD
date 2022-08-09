package transactions

import (
	"database/sql"

	"github.com/Xartyago/DDD/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.Transaction, error)
	Get(idToFind int) (domain.Transaction, error)
	Store(transactionCode, currency, emisor, receiver, transactionDate string, amount float64) (domain.Transaction, error)
	Update(idToUpdate int, transactionCode, currency, emisor, receiver, transactionDate string, amount float64) (domain.Transaction, error)
	PatchCode(idToPatch int, transactionCode string) (domain.Transaction, error)
	PatchAmount(idToPatch int, amount float64) (domain.Transaction, error)
	Delete(idToDelete int) (domain.Transaction, error)
}

type repository struct {
	dataBase *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		dataBase: db,
	}
}

// Get all Transactions
func (r *repository) GetAll() ([]domain.Transaction, error) {
	// vars
	var ts []domain.Transaction
	query := "SELECT * FROM transactions;"
	// Query
	rows, err := r.dataBase.Query(query)
	if err != nil {
		return []domain.Transaction{}, err
	}
	for rows.Next() {
		t := domain.Transaction{}
		rows.Scan(&t.Id, &t.TransactionCode, &t.Currency, &t.Emisor, &t.Receiver, &t.TransactionDate, &t.Amount)
		ts = append(ts, t)
	}
	return ts, nil
}

// Get one transaction
func (r *repository) Get(idToFind int) (domain.Transaction, error) {

	query := "SELECT transaction_code, currency, emisor, receiver, transaction_date, amount FROM transactions WHERE id = ?"
	// Get One
	var tsUpdated domain.Transaction
	tsUpdated.Id = idToFind
	row := r.dataBase.QueryRow(query, idToFind).Scan(&tsUpdated.TransactionCode, &tsUpdated.Currency, &tsUpdated.Emisor, &tsUpdated.Receiver, &tsUpdated.TransactionDate, &tsUpdated.Amount)
	if row != nil {
		return domain.Transaction{}, row
	}
	return tsUpdated, nil
}

// Create a new Transaction
func (r *repository) Store(transactionCode, currency, emisor, receiver, transactionDate string, amount float64) (domain.Transaction, error) {
	// Making Query
	query := "INSERT INTO transactions(transaction_code, currency, emisor, receiver, transaction_date, amount) VALUES (?,?,?,?,?,?)"
	stmt, err := r.dataBase.Prepare(query)
	if err != nil {
		return domain.Transaction{}, err
	}
	defer stmt.Close()
	// result
	result, err := stmt.Exec(&transactionCode, &currency, emisor, &receiver, &transactionDate, &amount)
	if err != nil {
		return domain.Transaction{}, err
	}
	insertedId, _ := result.LastInsertId()
	id := int(insertedId)
	// Create the new tction
	newTs := domain.Transaction{
		Id:              id,
		TransactionCode: transactionCode,
		Currency:        currency,
		Emisor:          emisor,
		Receiver:        receiver,
		TransactionDate: transactionDate,
		Amount:          amount,
	}
	return newTs, nil
}

// Update whole Transaction
func (r *repository) Update(idToUpdate int, transactionCode, currency, emisor, receiver, transactionDate string, amount float64) (domain.Transaction, error) {
	// Making query
	query := "UPDATE transactions SET transaction_code=?, currency=?, emisor=?, receiver=?, transaction_date=?, amount=? WHERE id=?"
	stmt, err := r.dataBase.Prepare(query)
	if err != nil {
		return domain.Transaction{}, err
	}
	defer stmt.Close()
	// result
	result, _ := stmt.Exec(transactionCode, currency, emisor, receiver, transactionDate, amount, idToUpdate)
	_, err = result.RowsAffected()
	if err != nil {
		return domain.Transaction{}, err
	}
	tsUpdated := domain.Transaction{
		Id:              idToUpdate,
		TransactionCode: transactionCode,
		Currency:        currency,
		Emisor:          emisor,
		Receiver:        receiver,
		TransactionDate: transactionDate,
		Amount:          amount,
	}
	return tsUpdated, nil
}

// Patch Transaction Code
func (r *repository) PatchCode(idToPatch int, transactionCode string) (domain.Transaction, error) {
	// Query
	query := "UPDATE transactions SET transaction_code=? WHERE id=?"
	stmt, err := r.dataBase.Prepare(query)
	if err != nil {
		return domain.Transaction{}, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(transactionCode, idToPatch)
	if err != nil {
		return domain.Transaction{}, err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return domain.Transaction{}, nil
	}
	return domain.Transaction{}, nil
}

// Patch Amount
func (r *repository) PatchAmount(idToPatch int, amount float64) (domain.Transaction, error) {
	// Query
	query := "UPDATE transactions SET amount=? WHERE id=?"
	stmt, err := r.dataBase.Prepare(query)
	if err != nil {
		return domain.Transaction{}, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(amount, idToPatch)
	if err != nil {
		return domain.Transaction{}, err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return domain.Transaction{}, nil
	}
	return domain.Transaction{}, nil
}

// Delete the transaction with the id specificed
func (r *repository) Delete(idToDelete int) (domain.Transaction, error) {
	query := "DELETE FROM transactions WHERE id=?"
	stmt, err := r.dataBase.Prepare(query)
	if err != nil {
		return domain.Transaction{}, nil
	}
	defer stmt.Close()
	res, _ := stmt.Exec(idToDelete)
	_, err = res.RowsAffected()
	if err != nil {
		return domain.Transaction{}, err
	}
	return domain.Transaction{}, nil
}
