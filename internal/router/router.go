package router

import (
	"product/internal/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	productHandler handler.ProductHandler,
) *gin.Engine {

	r := gin.New()

	// Middleware های عمومی
	r.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	api := r.Group("/api/v1")
	{
		products := api.Group("/products")
		{
			products.GET(
				"",
				productHandler.GetProducts(),
			)
		}
	}

	return r
}
