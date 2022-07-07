package transactions

type Repository interface {
	GetAll()
	Store()
}

type repository struct{}

func NewRepository() Repository {
	// return &repository{}
	return NewRepository()
}
