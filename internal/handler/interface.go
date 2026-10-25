package handler

import "github.com/gin-gonic/gin"

type ProductHandler interface {
	GetProducts() gin.HandlerFunc
	CreateProduct() gin.HandlerFunc
	UpdateProduct() gin.HandlerFunc
}
