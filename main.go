package main

import (
	"hacktiv8-assignment/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Products
	r.GET("/products", handlers.GetProducts)
	r.GET("/products/:id", handlers.GetProductByID)
	r.POST("/products", handlers.CreateProduct)
	r.PUT("/products/:id", handlers.UpdateProduct)
	r.DELETE("/products/:id", handlers.DeleteProduct)

	// Sources
	r.GET("/sources", handlers.GetSources)
	r.GET("/sources/:id", handlers.GetSourceByID)
	r.POST("/sources", handlers.CreateSource)
	r.PUT("/sources/:id", handlers.UpdateSource)
	r.DELETE("/sources/:id", handlers.DeleteSource)

	// Transactions
	r.GET("/transactions", handlers.GetTransactions)
	r.GET("/transactions/:id", handlers.GetTransactionByID)
	r.POST("/transactions", handlers.CreateTransaction)

	r.Run(":8080")
}
