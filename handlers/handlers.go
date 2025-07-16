package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"hacktiv8-assignment/models"
)

var products []models.Product
var sources []models.Source
var transactions []models.Transaction

// Helper: Response formatter
func response(ctx *gin.Context, status int, message string, data interface{}, err interface{}) {
	ctx.JSON(status, gin.H{
		"message": message,
		"data":    data,
		"error":   err,
	})
}

// ----------- PRODUCT HANDLERS -----------

func GetProducts(c *gin.Context) {
	response(c, 200, "List of products", products, nil)
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	for _, p := range products {
		if p.ID == id {
			response(c, 200, "Product found", p, nil)
			return
		}
	}
	response(c, 404, "Product not found", nil, "invalid_id")
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		response(c, 400, "Invalid input", nil, err.Error())
		return
	}
	if product.Price <= 0 || product.Stock < 0 {
		response(c, 400, "Invalid price or stock", nil, "Validation failed")
		return
	}
	product.ID = uuid.NewString()
	products = append(products, product)
	response(c, 201, "Product created", product, nil)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	for i, p := range products {
		if p.ID == id {
			var updated models.Product
			if err := c.ShouldBindJSON(&updated); err != nil {
				response(c, 400, "Invalid input", nil, err.Error())
				return
			}
			if updated.Price <= 0 || updated.Stock < 0 {
				response(c, 400, "Invalid price or stock", nil, "Validation failed")
				return
			}
			updated.ID = id
			products[i] = updated
			response(c, 200, "Product updated", updated, nil)
			return
		}
	}
	response(c, 404, "Product not found", nil, "invalid_id")
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			response(c, 200, "Product deleted", nil, nil)
			return
		}
	}
	response(c, 404, "Product not found", nil, "invalid_id")
}

// ----------- SOURCE HANDLERS -----------

func GetSources(c *gin.Context) {
	response(c, 200, "List of sources", sources, nil)
}

func GetSourceByID(c *gin.Context) {
	id := c.Param("id")
	for _, s := range sources {
		if s.ID == id {
			response(c, 200, "Source found", s, nil)
			return
		}
	}
	response(c, 404, "Source not found", nil, "invalid_id")
}

func CreateSource(c *gin.Context) {
	var source models.Source
	if err := c.ShouldBindJSON(&source); err != nil {
		response(c, 400, "Invalid input", nil, err.Error())
		return
	}
	source.ID = uuid.NewString()
	sources = append(sources, source)
	response(c, 201, "Source created", source, nil)
}

func UpdateSource(c *gin.Context) {
	id := c.Param("id")
	for i, s := range sources {
		if s.ID == id {
			var updated models.Source
			if err := c.ShouldBindJSON(&updated); err != nil {
				response(c, 400, "Invalid input", nil, err.Error())
				return
			}
			updated.ID = id
			sources[i] = updated
			response(c, 200, "Source updated", updated, nil)
			return
		}
	}
	response(c, 404, "Source not found", nil, "invalid_id")
}

func DeleteSource(c *gin.Context) {
	id := c.Param("id")
	for i, s := range sources {
		if s.ID == id {
			sources = append(sources[:i], sources[i+1:]...)
			response(c, 200, "Source deleted", nil, nil)
			return
		}
	}
	response(c, 404, "Source not found", nil, "invalid_id")
}

// ----------- TRANSACTION HANDLERS -----------

func GetTransactions(c *gin.Context) {
	response(c, 200, "List of transactions", transactions, nil)
}

func GetTransactionByID(c *gin.Context) {
	id := c.Param("id")
	for _, t := range transactions {
		if t.ID == id {
			response(c, 200, "Transaction found", t, nil)
			return
		}
	}
	response(c, 404, "Transaction not found", nil, "invalid_id")
}

func CreateTransaction(c *gin.Context) {
	var trx models.Transaction
	if err := c.ShouldBindJSON(&trx); err != nil {
		response(c, 400, "Invalid input", nil, err.Error())
		return
	}

	for i, p := range products {
		if p.ID == trx.ProductID {
			if trx.Quantity <= 0 || trx.Quantity > p.Stock {
				response(c, 400, "Invalid quantity", nil, "Stock not enough or invalid quantity")
				return
			}
			p.Stock -= trx.Quantity
			products[i] = p
			trx.ID = uuid.NewString()
			trx.Total = float64(trx.Quantity) * p.Price
			transactions = append(transactions, trx)
			response(c, 201, "Transaction success", trx, nil)
			return
		}
	}
	response(c, 404, "Product not found", nil, "invalid_product_id")
}
