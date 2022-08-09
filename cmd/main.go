package main

import (
	"log"
	"os"

	"github.com/Xartyago/DDD/cmd/handler"
	"github.com/Xartyago/DDD/db"
	"github.com/Xartyago/DDD/docs"
	"github.com/Xartyago/DDD/internal/transactions"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Transactions.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	db, err := db.NewConnection()
	if err != nil {
		log.Fatal("Error in the connection")
	}
	repository := transactions.NewRepository(db)
	service := transactions.NewService(repository)
	ts := handler.NewTransaction(service)
	// Routes
	r := gin.Default()

	//Swagger
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pr := r.Group("/transactions")
	pr.Use(TokenAuthMiddleware())
	pr.POST("/", ts.Store())
	pr.GET("/", ts.GetAll())
	pr.GET("/:id", ts.Get())
	pr.PUT("/:id", ts.Update())
	pr.PATCH("/code/:id", ts.PatchCode())
	pr.PATCH("/amount/:id", ts.PatchAmount())
	pr.DELETE("/:id", ts.Delete())
	r.Run(":8080")
}

func TokenAuthMiddleware() gin.HandlerFunc {
	envToken := os.Getenv("TOKEN")
	if envToken == "" {
		log.Fatal("doesnt find the .env var file")
	}
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"msg": "Token in headers not found"})
			return
		}
		if token != envToken {
			ctx.AbortWithStatusJSON(401, gin.H{"msg": "the token is not valid"})
			return
		}
	}
}
