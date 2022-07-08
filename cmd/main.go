package main

import (
	"github.com/Xartyago/DDD/cmd/handler"
	"github.com/Xartyago/DDD/internal/transactions"
	"github.com/gin-gonic/gin"
)

func main() {
	repository := transactions.NewRepository()
	service := transactions.NewService(repository)
	ts := handler.NewTransaction(service)

	r := gin.Default()
	pr := r.Group("/transactions")
	pr.POST("/", ts.Store())
	pr.GET("/", ts.GetAll())
	r.Run(":3001")
}
