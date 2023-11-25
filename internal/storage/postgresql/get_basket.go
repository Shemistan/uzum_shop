package postgresql

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/Shemistan/uzum_shop/internal/models"
)

func (r *repo) GetBasket(ctx context.Context, userId string) ([]*models.Basket, error) {
	var res []*models.Basket

	builder := sq.Select("id", "user_id", "product_id", "count").
		From("basket").
		Where(sq.Eq{"user_id": userId}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	rows, err := builder.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id, productId int64
		var userId string
		var count int32

		if err = rows.Scan(&id, &userId, &productId, &count); err != nil {
			return nil, err
		}

		res = append(res, &models.Basket{
			ID:        id,
			UserID:    userId,
			ProductID: productId,
			Count:     count,
		})
	}

	return res, nil
}
