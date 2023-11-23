package handlers

import (
	"BE-Golang-Coffee-Shop/internal/models"
	"BE-Golang-Coffee-Shop/internal/repositories"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerProduct struct {
	*repositories.ProductRepository
}

func InitializeProductHandler(r *repositories.ProductRepository) *HandlerProduct {
	return &HandlerProduct{r}
}

func (h *HandlerProduct) GetAllProduct(ctx *gin.Context) {
	name, search_name := ctx.GetQuery("name")

	if search_name {
		result, err := h.RepositorySearchProduct(name)

		if len(result) == 0 {
			// log.Println(err)
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "data not found",
			})
			return
		}
		if err != nil {
			// log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data":    result,
			"message": "search product success",
		})
		return
	}

	result, err := h.RepositoryGetAllProduct()
	if err != nil {
		// log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": "get all product success",
	})
}

func (h *HandlerProduct) GetProductID(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.RepositoryGetProductID(id)

	if len(result) == 0 {
		// log.Println(err)
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
		"message": "get product id success",
	})
}

func (h *HandlerProduct) CreateProduct(ctx *gin.Context) {
	var body models.ProductModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	result, err := h.RepositoryCreateProduct(&body)
	if err != nil {
		// log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": "create product success",
	})
}

func (h *HandlerProduct) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	var body models.ProductModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	result, err := h.RepositoryUpdateProduct(&body, id)
	if err != nil {
		// log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": "update product success",
	})
}

func (h *HandlerProduct) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	var body models.ProductModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	result, err := h.RepositoryDeleteProduct(&body, id)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": "delete product id: " + id + " success",
	})
}
