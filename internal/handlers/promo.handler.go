package handlers

import (
	"BE-Golang-Coffee-Shop/internal/models"
	"BE-Golang-Coffee-Shop/internal/repositories"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerPromo struct {
	*repositories.PromoRepository
}

func InitializePromoHandler(r *repositories.PromoRepository) *HandlerPromo {
	return &HandlerPromo{r}
}

func (h *HandlerPromo) GetAllPromo(ctx *gin.Context) {
	result, err := h.RepositoryGetAllPromo()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": "get all promo success",
	})
}

func (h *HandlerPromo) GetPromoID(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.RepositoryGetPromoID(id)

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
		"message": "get promo id success",
	})
}

func (h *HandlerPromo) CreatePromo(ctx *gin.Context) {
	var body models.PromoModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	result, err := h.RepositoryCreatePromo(&body)
	if err != nil {
		// log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": "create promo success",
	})
}

func (h *HandlerPromo) UpdatePromo(ctx *gin.Context) {
	id := ctx.Param("id")

	var body models.PromoModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	result, err := h.RepositoryUpdatePromo(&body, id)
	if err != nil {
		// log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": "update promo success",
	})
}

func (h *HandlerPromo) DeletePromo(ctx *gin.Context) {
	id := ctx.Param("id")
	var body models.PromoModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	result, err := h.RepositoryDeletePromo(&body, id)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": "delete promo id: " + id + " success",
	})
}
