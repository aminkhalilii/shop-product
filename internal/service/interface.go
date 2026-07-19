package service

import (
	"context"

	"product/internal/model"
)

type ProductService interface {
	GetProducts(
		ctx context.Context,
	) ([]model.Product, error)
}
