package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	/** Camada repository */
	productRepository := repository.NewProductRepository(dbConection)

	/** Camada casos de user */
	productUsecase := usecase.NewproductUsecase(productRepository)

	/** Camada controller */
	productController := controller.NewProductCrontroller(productUsecase)

	/** rotas */
	server.GET("/products", productController.GetProducts)
	server.POST("/product", productController.CreateProduct)
	server.GET("/product/:productId", productController.GetProductById)

	server.Run(":9001")
}
