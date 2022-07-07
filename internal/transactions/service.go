package transactions

import "github.com/Xartyago/DDD/internal/domain"

type Service interface {
	GetAll() ([]domain.Transaction, error)
	Store(id int, transactionCode, currency, emisor, receiver, transactionDate string, amount float64) (domain.Transaction, error)
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

func (s *service) Store(id int, transactionCode, currency, emisor, receiver, transactionDate string, amount float64) (domain.Transaction, error) {
	newTs, err := s.repository.Store(
		id, transactionCode, transactionDate, currency, emisor, receiver, amount,
	)
	if err != nil {
		return domain.Transaction{}, err
	}
	ts = append(ts, newTs)
	return newTs, nil
}
