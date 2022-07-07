package handler

import (
	"github.com/Xartyago/DDD/internal/transactions"
	"github.com/gin-gonic/gin"
)

type request struct {
	Id              int     `json:"id"`
	TransactionCode string  `json:"transactioncode"`
	Currency        string  `json:"currency"`
	Emisor          string  `json:"emisor"`
	Receiver        string  `json:"receiver"`
	TransactionDate string  `json:"transactiondate"`
	Amount          float64 `json:"amount"`
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(s transactions.Service) *Transaction {
	return &Transaction{
		service: s,
	}
}

func (s *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		t, err := s.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, t)
	}
}

func (s *Transaction) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(400, gin.H{"msg": "Invalid Token"})
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		t, err := s.service.Store(req.Id, req.Currency, req.Emisor, req.Receiver, req.TransactionCode, req.TransactionDate, req.Amount)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, t)
	}
}
