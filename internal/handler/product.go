package handler

import (
	"net/http"

	"product/internal/dto/request"
	"product/internal/service"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	service service.ProductService
}

var _ ProductHandler = (*productHandler)(nil)

func NewProductHandler(
	service service.ProductService,
) ProductHandler {

	return &productHandler{
		service: service,
	}
}

func (h *productHandler) GetProducts() gin.HandlerFunc {

	return func(c *gin.Context) {

		products, err := h.service.GetProducts(
			c.Request.Context(),
		)

		if err != nil {

			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error": err.Error(),
				},
			)

			return
		}

		c.JSON(
			http.StatusOK,
			gin.H{
				"data": products,
			},
		)
	}
}
func (h *productHandler) CreateProduct() gin.HandlerFunc {

	return func(c *gin.Context) {

		var req request.CreateProduct

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		product, err := h.service.Create(
			c.Request.Context(),
			req,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"data": product,
		})
	}
}
