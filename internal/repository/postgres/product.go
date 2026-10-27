package postgres

import (
	"context"
	"errors"

	"product/internal/model"
	"product/internal/repository"

	"github.com/jackc/pgx/v5"
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

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// ✅ جدید
func (r *ProductRepository) FindByID(
	ctx context.Context,
	id int,
) (*model.Product, error) {

	query := `
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
		WHERE id = $1
		LIMIT 1
	`

	var product model.Product

	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
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
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repository.ErrProductNotFound
		}
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) Create(
	ctx context.Context,
	product *model.Product,
) error {

	query := `
	INSERT INTO products
	(
		name,
		slug,
		description,
		price,
		stock
	)
	VALUES
	(
		$1,$2,$3,$4,$5
	)
	RETURNING id,created_at,updated_at
	`

	return r.db.QueryRow(
		ctx,
		query,
		product.Name,
		product.Slug,
		product.Description,
		product.Price,
		product.Stock,
	).Scan(
		&product.ID,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
}

func (r *ProductRepository) Update(
	ctx context.Context,
	product *model.Product,
) error {

	query := `
		UPDATE products
		SET
			name        = $1,
			slug        = $2,
			description = $3,
			price       = $4,
			stock       = $5,
			updated_at  = NOW()
		WHERE id = $6
		RETURNING updated_at
	`

	err := r.db.QueryRow(
		ctx,
		query,
		product.Name,
		product.Slug,
		product.Description,
		product.Price,
		product.Stock,
		product.ID,
	).Scan(&product.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return repository.ErrProductNotFound // ✅ اصلاح شد
		}
		return err
	}

	return nil
}

func (r *ProductRepository) Delete(
	ctx context.Context,
	id int,
) error {

	query := `
		DELETE FROM products
		WHERE id = $1
	`

	tag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return repository.ErrProductNotFound
	}

	return nil
}
