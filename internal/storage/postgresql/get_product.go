package postgresql

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/Shemistan/uzum_shop/internal/models"
)

func (r *repo) GetProduct(ctx context.Context, id int64) (*models.Product, error) {
	var res models.Product

	builder := sq.Select("id", "name", "description", "price", "count").
		From("products").
		Where(sq.Eq{"id": id}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	err := builder.QueryRowContext(ctx).Scan(&res.ID, &res.Name, &res.Description, &res.Price, &res.Count)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
