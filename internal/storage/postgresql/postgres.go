package postgresql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/Shemistan/uzum_shop/internal/storage"
)

type repo struct {
	db *sqlx.DB
}

func NewRepoPostgres(db *sqlx.DB) storage.IStorage {
	return &repo{db: db}
}

func (r *repo) getProductCount(ctx context.Context, id int64) (int64, error) {
	var count int64

	query := sq.Select("count").
		From("products").
		Where(sq.Eq{"id": id}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	err := query.QueryRowContext(ctx).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
