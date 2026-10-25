package handler

import (
	"net/http"
	"strconv"

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
func (h *productHandler) UpdateProduct() gin.HandlerFunc {

	return func(c *gin.Context) {

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil || id <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid product id",
			})
			return
		}

		var req request.UpdateProduct

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		product, err := h.service.Update(
			c.Request.Context(),
			id,
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
