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
	pr.PUT("/", ts.Update())
	pr.PATCH("/code/:id", ts.PatchCode())
	pr.PATCH("/amount/:id", ts.PatchAmount())
	pr.DELETE("/", ts.Delete())
	r.Run(":3001")
}
