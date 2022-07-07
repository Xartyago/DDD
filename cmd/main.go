package main

import (
	"fmt"

	"github.com/Xartyago/DDD/internal/transactions"
)

func main() {
	repository := transactions.NewRepository()
	service := transactions.NewService(repository)
	fmt.Println(service)
}
