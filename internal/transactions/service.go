package transactions

import "github.com/Xartyago/DDD/internal/domain"

type Service interface {
	GetAll() ([]domain.Transaction, error)
	Store(transactionCode, currency, emisor, receiver, transactionDate string, amount float64) (domain.Transaction, error)
	Update(idToUpdate int, transactionCode, currency, emisor, receiver, transactionDate string, amount float64) (domain.Transaction, error)
	PatchCode(idToPatch int, transactionCode string) (domain.Transaction, error)
	PatchAmount(idToPatch int, amount float64) (domain.Transaction, error)
	Delete(idToDelete int) (domain.Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
func (s *service) GetAll() ([]domain.Transaction, error) {
	ts, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func (s *service) Store(transactionCode, currency, emisor, receiver, transactionDate string, amount float64) (domain.Transaction, error) {
	newTs, err := s.repository.Store(
		transactionCode,
		currency,
		emisor,
		receiver,
		transactionDate,
		amount,
	)
	if err != nil {
		return domain.Transaction{}, err
	}
	ts = append(ts, newTs)
	return newTs, nil
}
func (s *service) Update(idToUpdate int, transactionCode, currency, emisor, receiver, transactionDate string, amount float64) (domain.Transaction, error) {
	return s.repository.Update(idToUpdate, transactionCode, currency, emisor, receiver, transactionDate, amount)
}

func (s *service) PatchCode(idToPatch int, transactionCode string) (domain.Transaction, error) {
	return s.repository.PatchCode(idToPatch, transactionCode)
}
func (s *service) PatchAmount(idToPatch int, amount float64) (domain.Transaction, error) {
	return s.repository.PatchAmount(idToPatch, amount)
}
func (s *service) Delete(idToDelete int) (domain.Transaction, error) {
	return s.repository.Delete(idToDelete)
}
