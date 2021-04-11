package servies

import (
	"cab-management-portal/app/models"
	"context"
)

func (s *Service) CreateCity(ctx context.Context, city *models.City) error {
	_, err := s.Sequel.NamedExecContext(
		ctx,
		`INSERT INTO `+models.CitiesTableName+
			` (name,is_active) VALUES `+
			`(:name,:is_active)`,
		city,
	)
	return err
}

func (s *Service) GetCity(ctx context.Context, id int) (*models.City, error) {
	city := models.City{}
	err := s.Sequel.GetContext(ctx, &city, "SELECT * FROM " +
		models.CitiesTableName +
		" WHERE id=?",id)
	if err != nil {
		return nil, err
	}
	return &city, nil
}

func (s *Service) GetAllCities(ctx context.Context) ([]models.City, error) {
	var cities []models.City
	err := s.Sequel.SelectContext(ctx, &cities, "SELECT * FROM "+
		models.CitiesTableName)
	if err != nil {
		return nil, err
	}
	if cities == nil {
		cities = []models.City{}
	}
	return cities, nil
}