package main

import (
	"BE-Golang-Coffee-Shop/internal/routers"
	"BE-Golang-Coffee-Shop/pkg"
	"log"
)

func main() {
	database, err := pkg.PostgreSQLDB()
	if (err) != nil {
		log.Fatal(err)
	}

	routers := routers.New(database)
	server := pkg.Server(routers)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

// func main() {
// 	r := gin.Default()
// 	r.GET("/", func(ctx *gin.Context) {
// 		ctx.JSON(http.StatusOK, gin.H{
// 			"message": "success",
// 		})
// 	})
// 	r.GET("/product", GetAllProduct)

// 	r.Run("localhost:8000")
// }

// type ProductQuery struct {
// 	Search string `form:"search"`
// 	Limit  string `form:"limit"`
// }

// func GetAllProduct(ctx *gin.Context) {
// 	// search := ctx.Query("search")
// 	// limit := ctx.Query("limit")

// 	var query ProductQuery
// 	if err := ctx.ShouldBind(&query); err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	search := query.Search
// 	limit := query.Limit
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"search": search,
// 		"limit":  limit,
// 	})
// }
