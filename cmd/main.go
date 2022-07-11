package main

import (
	"log"

	"github.com/Xartyago/DDD/cmd/handler"
	"github.com/Xartyago/DDD/internal/transactions"
	"github.com/Xartyago/DDD/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Get the enviroment vars
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("error in load .env file")
	}
	//Initialize the whole server
	db := store.NewStore("../transactions.json")
	repository := transactions.NewRepository(db)
	service := transactions.NewService(repository)
	ts := handler.NewTransaction(service)

	// Routes
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
