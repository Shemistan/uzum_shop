package postgresql

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *storage) CancelOrderStorage(ctx context.Context, req *pb.CancelOrder_Request) (*empty.Empty, error) {
	q := sq.Delete("orders").
		Where(sq.Eq{"id": req.OrderId}).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar)

	_, err := q.ExecContext(ctx)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
