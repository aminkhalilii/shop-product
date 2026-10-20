package service

import (
	"context"

	"product/internal/dto/request"
	"product/internal/dto/response"
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
) ([]response.Product, error) {

	products, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]response.Product, 0, len(products))
	for _, p := range products {
		result = append(result, toProductResponse(p))
	}

	return result, nil
}

func (s *productService) Create(
	ctx context.Context,
	req request.CreateProduct,
) (*response.Product, error) {

	product := &model.Product{
		Name:        req.Name,
		Slug:        req.Slug,
		Description: &req.Description,
		Price:       int64(req.Price),
		Stock:       req.Stock,
	}

	if err := s.repo.Create(ctx, product); err != nil {
		return nil, err
	}

	resp := toProductResponse(*product)
	return &resp, nil
}

func toProductResponse(p model.Product) response.Product {
	desc := ""
	if p.Description != nil {
		desc = *p.Description
	}

	return response.Product{
		ID:          p.ID,
		Name:        p.Name,
		Slug:        p.Slug,
		Description: desc,
		Price:       float64(p.Price),
		Stock:       p.Stock,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}
