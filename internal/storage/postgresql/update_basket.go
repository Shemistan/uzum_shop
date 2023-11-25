package postgresql

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/Shemistan/uzum_shop/internal/models"
)

func (r *repo) UpdateBasket(ctx context.Context, req *models.UpdateBasketRequest) error {
	q := sq.Update("basket").SetMap(map[string]interface{}{
		"product_id": req.ProductID,
		"count":      req.Count,
	}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	_, err := q.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
