package repository

import (
	"context"

	"product/internal/model"
)

type ProductRepository interface {
	FindAll(
		ctx context.Context,
	) ([]model.Product, error)
}
