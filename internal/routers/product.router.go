package routers

import (
	"BE-Golang-Coffee-Shop/internal/handlers"
	"BE-Golang-Coffee-Shop/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RouterProduct(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/product")
	repository := repositories.InitializeProductRepository(db)

	handler := handlers.InitializeProductHandler(repository)

	route.GET("/", handler.GetAllProduct)
	route.GET("/:id", handler.GetProductID)
	// route.GET("/search", handler.GetAllProduct)
	route.POST("/", handler.CreateProduct)
	route.PATCH("/:id", handler.UpdateProduct)
	route.DELETE("/:id", handler.DeleteProduct)
}
