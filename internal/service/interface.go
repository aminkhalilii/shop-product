package service

import (
	"context"

	"product/internal/dto/request"
	"product/internal/dto/response"
)

type ProductService interface {
	GetProducts(
		ctx context.Context,
	) ([]response.Product, error)
	Create(
		ctx context.Context,
		req request.CreateProduct,
	) (*response.Product, error)
}
