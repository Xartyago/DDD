package transactions

import (
	"fmt"
	"testing"

	"github.com/Xartyago/DDD/internal/domain"
	"github.com/stretchr/testify/assert"
)

type MockStorage struct {
	ts            []domain.Transaction
	ReadWasCalled bool
}

func (fs *MockStorage) Read(data interface{}) error {
	fs.ReadWasCalled = true
	toFillTs := data.(*[]domain.Transaction)
	*toFillTs = fs.ts
	return nil
}

func (fs *MockStorage) Write(data interface{}) error {
	toFillMock := data.([]domain.Transaction)
	fmt.Println(toFillMock)
	fs.ts = toFillMock
	return nil
}
func TestUpdate(t *testing.T) {
	transactionsBefore := []domain.Transaction{
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
	transactionAfter := domain.Transaction{
		Id:              2,
		TransactionCode: "3ee3a293-b467-42b8-a9f7-f835fcd44c34",
		Currency:        "COP",
		Emisor:          "Cherie Tillman",
		Receiver:        "QUANTALIA",
		TransactionDate: "2020-09-24",
		Amount:          5000,
	}

	db := &MockStorage{ts: transactionsBefore}
	repo := NewRepository(db)

	result, _ := repo.PatchAmount(2, 5000)
	secondResult, err := repo.GetAll()

	assert.Equal(t, secondResult, transactionsBefore)
	assert.Nil(t, err)
	assert.True(t, db.ReadWasCalled)
	assert.Equal(t, result, transactionAfter)
}
