package service

import (
	"context"

	"product/internal/model"
	"product/internal/repository"
)

type productService struct {
	repo repository.ProductRepository
}

var _ ProductService = (*productService)(nil)

func NewProductService(
	repo repository.ProductRepository,
) ProductService {

	return &productService{
		repo: repo,
	}
}

func (s *productService) GetProducts(
	ctx context.Context,
) ([]model.Product, error) {

	return s.repo.FindAll(ctx)
}
