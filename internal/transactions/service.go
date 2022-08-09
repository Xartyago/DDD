package transactions

import "github.com/Xartyago/DDD/internal/domain"

type Service interface {
	GetAll() ([]domain.Transaction, error)
	Get(idToFind int) (domain.Transaction, error)
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

func (s *service) Get(idToFind int) (domain.Transaction, error) {
	ts, err := s.repository.Get(idToFind)
	if err != nil {
		return domain.Transaction{}, nil
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
	return newTs, nil
}

func (s *service) Update(idToUpdate int, transactionCode, currency, emisor, receiver, transactionDate string, amount float64) (domain.Transaction, error) {
	tsUpdated, err := s.repository.Update(idToUpdate, transactionCode, currency, emisor, receiver, transactionDate, amount)
	if err != nil {
		return domain.Transaction{}, err
	}
	return tsUpdated, nil
}

func (s *service) PatchCode(idToPatch int, transactionCode string) (domain.Transaction, error) {
	_, err := s.repository.PatchCode(idToPatch, transactionCode)
	if err != nil {
		return domain.Transaction{}, nil
	}
	ts, err := s.repository.Get(idToPatch)
	if err != nil {
		return domain.Transaction{}, nil
	}
	return ts, nil
}

func (s *service) PatchAmount(idToPatch int, amount float64) (domain.Transaction, error) {
	_, err := s.repository.PatchAmount(idToPatch, amount)
	if err != nil {
		return domain.Transaction{}, nil
	}
	ts, err := s.repository.Get(idToPatch)
	if err != nil {
		return domain.Transaction{}, nil
	}
	return ts, nil
}

func (s *service) Delete(idToDelete int) (domain.Transaction, error) {
	return s.repository.Delete(idToDelete)
}
