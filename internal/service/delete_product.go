package service

import "context"

func (s service) DeleteProduct(ctx context.Context, id int64) error {
	// check user
	_, err := s.userId(ctx)
	if err != nil {
		return err
	}

	err = s.storage.DeleteProduct(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
