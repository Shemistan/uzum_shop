package postgresql

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/Shemistan/uzum_shop/internal/storage"
)

type repo struct {
	db *sqlx.DB
}

func NewRepoPostgres(db *sqlx.DB) storage.IStorage {
	return &repo{db: db}
}

func (r *repo) getProductCount(ctx context.Context) (int64, error) {
	var count int64

	query := squirrel.Select("COUNT(*)").From("products")

	err := query.RunWith(r.db).QueryRowContext(ctx).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
