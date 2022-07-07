package transactions

import (
	"git@github.com:Xartyago/DDD.git/domain"
)

type Repository interface {
	GetAll() domain.
	Store()
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}
