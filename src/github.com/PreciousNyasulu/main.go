package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type product struct {
	ID          string `json:"id"`
	ProductName string `json:"productname"`
	Barcode     string `json:"barcode"`
	Category    string `json:"category"`
	Brand       string `json:"brand"`
	Quantity       int `json:"quantity"`
}

var products = []product{
	{ID: "1", ProductName: "Product 1", Barcode: "123456789", Category: "Category A", Brand: "Brand X",Quantity: 10},
	{ID: "2", ProductName: "Product 2", Barcode: "987654321", Category: "Category B", Brand: "Brand Y",Quantity: 10},
}

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

func createProduct(c *gin.Context) {
	var newProduct product
	if err := c.BindJSON(&newProduct); err != nil {
		
		return 
	}

	products = append(products, newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)
}

func productById(c *gin.Context) {
	id := c.Param("id")
	product, err := getProductById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Product not found"})
		return 
	}
	c.IndentedJSON(http.StatusOK,product)
}

func getProductById(id string) (*product, error) {
	for i, b := range products {
		if b.ID == id {
			return &products[i], nil
		}
	}
	return nil, errors.New("Product not found")
}

func productTender(c *gin.Context){
	id,ok := c.GetQuery("id")

	if !ok  {
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Missing id query parameter"})
		return 
	}

	product,err := getProductById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Product not found."})
		return
	}

	if product.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest,gin.H{"Message":"Product not available"})
		return
	}

	product.Quantity -= 1
	c.IndentedJSON(http.StatusOK,product)
}

func returnProduct(c *gin.Context){
	id,ok := c.GetQuery("id")

	if !ok  {
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Missing id query parameter"})
		return 
	}

	product,err := getProductById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Product not found."})
		return
	}

	product.Quantity += 1
	c.IndentedJSON(http.StatusOK,product)
}



func main() {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.GET("/products/:id", productById)
	router.POST("/createproduct", createProduct)
	router.PATCH("/product/tender", productTender)
	router.PATCH("/product/return", returnProduct)
	router.Run("localhost:8000")
}
