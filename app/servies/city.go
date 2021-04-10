package servies

import (
	"cab-management-portal/app/models"
	"context"
)

func (s *Services) CreateCity(ctx context.Context, city *models.City) error {
	_, err := s.Sequel.NamedExecContext(
		ctx,
		`INSERT INTO `+models.CitiesTableName+
			` (name,is_active,last_updated_by) VALUES `+
			`(:name,:is_active,:last_updated_by)`,
		city,
	)
	return err
}
