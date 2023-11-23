package routers

import (
	"BE-Golang-Coffee-Shop/internal/handlers"
	"BE-Golang-Coffee-Shop/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RouterUser(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/user")
	repository := repositories.InitializeUserRepository(db)

	handler := handlers.InitializeUserHandler(repository)

	route.GET("/", handler.GetAllUser)
	route.GET("/:id", handler.GetUserID)
	route.POST("/", handler.CreateUser)
	route.PATCH("/:id", handler.UpdateUser)
	route.DELETE("/:id", handler.DeleteUser)
}
