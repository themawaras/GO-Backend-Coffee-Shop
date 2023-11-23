package handlers

import (
	"BE-Golang-Coffee-Shop/internal/models"
	"BE-Golang-Coffee-Shop/internal/repositories"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerOrder struct {
	*repositories.OrderRepository
}

func InitializeOrderHandler(r *repositories.OrderRepository) *HandlerOrder {
	return &HandlerOrder{r}
}

func (h *HandlerOrder) GetAllOrder(ctx *gin.Context) {
	result, err := h.RepositoryGetAllOrder()
	if err != nil {
		// log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": "get all order success",
	})
}

func (h *HandlerOrder) GetOrderID(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.RepositoryGetOrderID(id)

	if len(result) == 0 {
		log.Println(err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "data not found",
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": "get order id success",
	})
}

func (h *HandlerOrder) CreateOrder(ctx *gin.Context) {
	var body models.OrderModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	result, err := h.RepositoryCreateOrder(&body)
	if err != nil {
		// log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": "create order success",
	})
}

func (h *HandlerOrder) UpdateOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	var body models.OrderModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	result, err := h.RepositoryUpdateOrder(&body, id)
	if err != nil {
		// log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": "update order: " + id + " success",
	})
}

func (h *HandlerOrder) DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	var body models.OrderModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	result, err := h.RepositoryDeleteOrder(&body, id)
	if err != nil {
		// log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": "delete order id: " + id + " success",
	})
}
