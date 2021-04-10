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

func (s *Services) GetCity(ctx context.Context, id int) (*models.City, error) {
	city := models.City{}
	err := s.Sequel.GetContext(ctx, &city, "SELECT * FROM " +
		models.CitiesTableName +
		" WHERE id=?",id)
	if err != nil {
		return nil, err
	}
	return &city, nil
}