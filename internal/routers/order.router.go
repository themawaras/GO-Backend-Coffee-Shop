package routers

import (
	"BE-Golang-Coffee-Shop/internal/handlers"
	"BE-Golang-Coffee-Shop/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RouterOrder(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/order")
	repository := repositories.InitializeOrderRepository(db)

	handler := handlers.InitializeOrderHandler(repository)

	route.GET("/", handler.GetAllOrder)
	route.GET("/:id", handler.GetOrderID)
	route.POST("/", handler.CreateOrder)
	route.PATCH("/:id", handler.UpdateOrder)
	route.DELETE("/:id", handler.DeleteOrder)
}
