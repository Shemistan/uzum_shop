package postgresql

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
)

func (s *storage) GetBasketStorage(ctx context.Context, userId int) (*pb.GetBasket_Response, error) {
	var res []*pb.BasketItem

	builder := sq.Select("id", "user_id", "product_id", "amount").
		From("basket").
		Where(sq.Eq{"user_id": userId}).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar)

	rows, err := builder.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id uint32
		var user_id uint32
		var product_id uint32
		var count uint32

		if err = rows.Scan(&id, &user_id, &product_id, &count); err != nil {
			return nil, err
		}

		res = append(res, &pb.BasketItem{
			Id:        id,
			UserId:    user_id,
			ProductId: product_id,
			Count:     count,
		})
	}
	respond := &pb.GetBasket_Response{
		AllBasket: res,
	}

	return respond, nil
}
