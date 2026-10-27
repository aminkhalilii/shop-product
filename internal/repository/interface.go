package repository

import (
	"context"
	"errors"

	"product/internal/model"
)

var ErrProductNotFound = errors.New("product not found")

type ProductRepository interface {
	FindAll(
		ctx context.Context,
	) ([]model.Product, error)
	FindByID(
		ctx context.Context,
		id int,
	) (*model.Product, error) // ✅ اضافه شد
	Create(
		ctx context.Context,
		product *model.Product,
	) error
	Update(
		ctx context.Context,
		product *model.Product,
	) error
	Delete(ctx context.Context, id int) error
}
