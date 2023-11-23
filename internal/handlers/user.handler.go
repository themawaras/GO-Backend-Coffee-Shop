package handlers

import (
	"BE-Golang-Coffee-Shop/internal/models"
	"BE-Golang-Coffee-Shop/internal/repositories"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	*repositories.UserRepository
}

func InitializeUserHandler(r *repositories.UserRepository) *HandlerUser {
	return &HandlerUser{r}
}

func (h *HandlerUser) GetAllUser(ctx *gin.Context) {
	result, err := h.RepositoryGetAllUser()
	if err != nil {
		// log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": "get all user success",
	})
}

func (h *HandlerUser) GetUserID(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.RepositoryGetUserID(id)

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
		"message": "get product id success",
	})
}

func (h *HandlerUser) CreateUser(ctx *gin.Context) {
	var body models.UserModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	result, err := h.RepositoryCreateUser(&body)
	if err != nil {
		// log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": "create user success",
	})
}

func (h *HandlerUser) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var body models.UserModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	result, err := h.RepositoryUpdateUser(&body, id)
	if err != nil {
		// log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": "update user: " + id + " success",
	})
}

func (h *HandlerUser) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var body models.UserModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	result, err := h.RepositoryDeleteUser(&body, id)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": "delete user id: " + id + " success",
	})
}
