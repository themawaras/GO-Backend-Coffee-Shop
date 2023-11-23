package routers

import (
	"BE-Golang-Coffee-Shop/internal/handlers"
	"BE-Golang-Coffee-Shop/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RouterPromo(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/promo")
	repository := repositories.InitializePromoRepository(db)

	handler := handlers.InitializePromoHandler(repository)

	route.GET("/", handler.GetAllPromo)
	route.GET("/:id", handler.GetPromoID)
	route.POST("/", handler.CreatePromo)
	route.PATCH("/:id", handler.UpdatePromo)
	route.DELETE("/:id", handler.DeletePromo)
}
