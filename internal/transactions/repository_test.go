package transactions

import (
	"testing"

	"github.com/Xartyago/DDD/internal/domain"
	"github.com/stretchr/testify/assert"
)

type StubStore struct{}

func (fs *StubStore) Read(data interface{}) error {
	f := data.(*[]domain.Transaction)
	*f = []domain.Transaction{
		{
			Id:              1,
			TransactionCode: "37b7dcd4-ad7b-4e9d-bfd1-2d9e7cdf94ad",
			Currency:        "ARS",
			Emisor:          "Anna George",
			Receiver:        "ISOPOP",
			TransactionDate: "2020-10-12",
			Amount:          335.9562,
		},
		{
			Id:              2,
			TransactionCode: "3ee3a293-b467-42b8-a9f7-f835fcd44c34",
			Currency:        "COP",
			Emisor:          "Cherie Tillman",
			Receiver:        "QUANTALIA",
			TransactionDate: "2020-09-24",
			Amount:          99.9797,
		},
	}
	return nil
}

func (fs *StubStore) Write(data interface{}) error {
	return nil
}

// Stub Test
func TestRead(t *testing.T) {
	stub := &StubStore{}
	repo := NewRepository(stub)
	expectedResult := []domain.Transaction{
		{
			Id:              1,
			TransactionCode: "37b7dcd4-ad7b-4e9d-bfd1-2d9e7cdf94ad",
			Currency:        "ARS",
			Emisor:          "Anna George",
			Receiver:        "ISOPOP",
			TransactionDate: "2020-10-12",
			Amount:          335.9562,
		},
		{
			Id:              2,
			TransactionCode: "3ee3a293-b467-42b8-a9f7-f835fcd44c34",
			Currency:        "COP",
			Emisor:          "Cherie Tillman",
			Receiver:        "QUANTALIA",
			TransactionDate: "2020-09-24",
			Amount:          99.9797,
		},
	}
	ts, err := repo.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, ts, expectedResult)
}
