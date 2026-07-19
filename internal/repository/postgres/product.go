package postgres

import (
	"context"

	"product/internal/model"
	"product/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	db *pgxpool.Pool
}

var _ repository.ProductRepository = (*ProductRepository)(nil)

func NewProductRepository(
	db *pgxpool.Pool,
) *ProductRepository {

	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) FindAll(
	ctx context.Context,
) ([]model.Product, error) {

	rows, err := r.db.Query(
		ctx,
		`
		SELECT
			id,
			name,
			slug,
			description,
			price,
			stock,
			created_at,
			updated_at
		FROM products
		ORDER BY id DESC
		`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := make([]model.Product, 0)

	for rows.Next() {

		var product model.Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Slug,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.CreatedAt,
			&product.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		products = append(
			products,
			product,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
