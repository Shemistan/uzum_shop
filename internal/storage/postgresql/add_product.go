package postgresql

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"

	"github.com/Shemistan/uzum_shop/internal/models"
)

func (r *repo) AddProduct(ctx context.Context, req *models.Basket) error {
	count, countErr := r.getProductCount(ctx)
	if countErr != nil {
		return countErr
	}
	if count <= 0 {
		return errors.New("empty product")
	}

	builder := sq.Insert("basket").
		Columns("user_id", "product_id", "count").
		Values(req.UserID, req.ProductID, req.Count).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	_, err := builder.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
