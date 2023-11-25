package postgresql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (r *repo) DeleteProduct(ctx context.Context, id int64) error {
	q := sq.Delete("basket").
		Where(sq.Eq{"id": id}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	_, err := q.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
