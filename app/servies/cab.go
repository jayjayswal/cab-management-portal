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

func (s *Services) UpdateCabCity(ctx context.Context, cab *models.Cab) error {
	_, err := s.Sequel.NamedExecContext(
		ctx,
		"UPDATE "+models.CabsTableName+
			"SET current_city_id=:current_city_id,"+
			"WHERE id=:id",
		cab,
	)
	return err
}

func (s *Services) GetCab(ctx context.Context, id int) (*models.Cab, error) {
	cab := models.Cab{}
	err := s.Sequel.GetContext(ctx, &cab, "SELECT * FROM "+
		models.CabsTableName+
		" WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	return &cab, nil
}

func (s *Services) GetMostIdleCabOfCity(ctx context.Context, cityId int) (*models.Cab, error) {
	cab := models.Cab{}
	err := s.Sequel.GetContext(ctx, &cab,
		"SELECT *, IFNULL(last_ride_end_time, created) as last_ride_time "+
			"FROM ? "+
			"WHERE is_active = 1 AND current_city_id = ? "+
			"WHERE AND current_state = ? "+
			"ORDER BY last_ride_time "+
			"LIMIT 1", models.CabsTableName, models.CabIdleState, cityId)
	if err != nil {
		return nil, err
	}
	return &cab, nil
}
