package handler

import (
	"strconv"

	"github.com/Xartyago/DDD/internal/transactions"
	"github.com/gin-gonic/gin"
)

type request struct {
	TransactionCode string  `json:"transaction_code"`
	Currency        string  `json:"currency"`
	Emisor          string  `json:"emisor"`
	Receiver        string  `json:"receiver"`
	TransactionDate string  `json:"transaction_date"`
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

// ListTransactions godoc
// @Summary List transactions
// @Tags Transactions
// @Description get transactions
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /transactions [get]
func (s *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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
func (s *Transaction) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idToFind, _ := strconv.Atoi(ctx.Param("id"))
		t, err := s.service.Get(idToFind)
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
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		if req.TransactionCode == "" {
			ctx.JSON(400, gin.H{"error": "The transaction code is requerid"})
			return
		}
		if req.Currency == "" {
			ctx.JSON(400, gin.H{"error": "The currency is requerid"})
			return
		}
		if req.Emisor == "" {
			ctx.JSON(400, gin.H{"error": "The emisor is requerid"})
			return
		}
		if req.Receiver == "" {
			ctx.JSON(400, gin.H{"error": "The receiver is requerid"})
			return
		}
		if req.TransactionDate == "" {
			ctx.JSON(400, gin.H{"error": "The transaction date is requerid"})
			return
		}
		if req.Amount < 0 {
			ctx.JSON(400, gin.H{"error": "The amount cannt be empty"})
			return
		}
		t, err := s.service.Store(req.TransactionCode, req.Currency, req.Emisor, req.Receiver, req.TransactionDate, req.Amount)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, t)
	}
}

func (s *Transaction) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		idToUpdate, _ := strconv.Atoi(ctx.Param("id"))
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.TransactionCode == "" {
			ctx.JSON(400, gin.H{"error": "The transaction code is requerid"})
			return
		}
		if req.Currency == "" {
			ctx.JSON(400, gin.H{"error": "The currency is requerid"})
			return
		}
		if req.Emisor == "" {
			ctx.JSON(400, gin.H{"error": "The emisor is requerid"})
			return
		}
		if req.Receiver == "" {
			ctx.JSON(400, gin.H{"error": "The receiver is requerid"})
			return
		}
		if req.TransactionDate == "" {
			ctx.JSON(400, gin.H{"error": "The transaction date is requerid"})
			return
		}
		if req.Amount < 0 {
			ctx.JSON(400, gin.H{"error": "The amount cannt be empty"})
			return
		}
		ts, err := s.service.Update(idToUpdate, req.TransactionCode, req.Currency, req.Emisor, req.Receiver, req.TransactionCode, req.Amount)
		if err != nil {
			ctx.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		ctx.JSON(200, ts)
	}
}

func (s *Transaction) PatchCode() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		idToPatch, _ := strconv.Atoi(ctx.Param("id"))
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.TransactionCode == "" {
			ctx.JSON(400, gin.H{"error": "The transaction code cant be empty"})
		}
		ts, err := s.service.PatchCode(idToPatch, req.TransactionCode)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, ts)
	}
}

func (s *Transaction) PatchAmount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		idToPatch, _ := strconv.Atoi(ctx.Param("id"))
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.Amount < 0 {
			ctx.JSON(400, gin.H{"error": "The transaction code cant 0"})
		}
		ts, err := s.service.PatchAmount(idToPatch, req.Amount)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, ts)
	}
}

func (s *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idToDelete, _ := strconv.Atoi(ctx.Param("id"))
		ts, err := s.service.Delete(idToDelete)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"deleted": ts})
	}
}
