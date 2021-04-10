package servies

import (
	"cab-management-portal/app/models"
	"context"
)

func (s *Services) CreateCab(ctx context.Context, cab *models.Cab) error {
	_, err := s.Sequel.NamedExecContext(
		ctx,
		`INSERT INTO `+models.CabsTableName+
			` (current_state,current_city_id,is_active,last_updated_by,cab_number) VALUES `+
			`(:current_state,:current_city_id,:is_active,:last_updated_by,:cab_number)`,
		cab,
	)
	return err
}
