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

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//CAMADA DE REPOSITORY
	ProductRepository := repository.NewProductRepository(dbConnection)

	//CAMADA USECASE
	//INICIALIZAR O USECASE
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	//CAMADA DE CONTROLLERS
	//INICIALIZAR O CONTROLLER
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("product/:productId", ProductController.GetProductById)

	server.Run(":8000")
}
