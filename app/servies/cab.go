package servies

import (
	"cab-management-portal/app/models"
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
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

func (s *Services) GetCabActivities(ctx context.Context, id int) ([]models.CabAudit, error) {
	var cabActivities []models.CabAudit
	err := s.Sequel.SelectContext(ctx, &cabActivities, "SELECT * FROM "+
		models.CabsAuditTableName+
		" WHERE id=? ORDER BY updated DESC LIMIT 100", id)
	if err != nil {
		return nil, err
	}
	if cabActivities == nil {
		cabActivities = []models.CabAudit{}
	}
	return cabActivities, nil
}

func (s *Services) GetCabForUpdate(ctx context.Context, id int, tx *sqlx.Tx) (*models.Cab, error) {
	cab := models.Cab{}
	err := tx.GetContext(ctx, &cab, "SELECT * FROM "+
		models.CabsTableName+
		" WHERE id=? FOR UPDATE", id)
	if err != nil {
		return nil, err
	}
	return &cab, nil
}

func (s *Services) GetMostIdleCabOfCity(ctx context.Context, cityId int, tx *sqlx.Tx) ([]models.Cab, error) {
	var cab []models.Cab
	err := tx.SelectContext(ctx, &cab,
		"SELECT id, cab_number, current_state, current_city_id"+
			", is_active, last_updated_by, created, updated"+
			", IFNULL(last_ride_end_time, created) as last_ride_end_time "+
			"FROM "+models.CabsTableName+" "+
			"WHERE is_active = 1 AND current_city_id = ? "+
			"AND current_state = ? "+
			"ORDER BY last_ride_end_time "+
			"LIMIT 1 FOR UPDATE",  cityId, models.CabIdleState)
	if err != nil {
		return nil, err
	}
	return cab, nil
}

func (s *Services) UpdateCabCity(ctx context.Context, cab *models.Cab, tx *sqlx.Tx) error {
	query := "UPDATE "+models.CabsTableName+" "+
		"SET current_city_id=:current_city_id, last_updated_by=:last_updated_by "+
		"WHERE id=:id"
	var err error
	if tx != nil {
		_, err = tx.NamedExecContext(
			ctx, query, cab,
		)
	} else {
		_, err = s.Sequel.NamedExecContext(
			ctx, query, cab,
		)
	}
	return err
}

func (s *Services) UpdateCabState(ctx context.Context, cab *models.Cab, tx *sqlx.Tx) error {
	res, err := tx.NamedExecContext(
		ctx,
		"UPDATE "+models.CabsTableName+ " "+
			"SET current_state=:current_state, last_updated_by=:last_updated_by, "+
			"current_city_id=:current_city_id "+
			"WHERE id=:id",
		cab,
	)
	if err == nil {
		totalRows, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if totalRows == 0 {
			return errors.New("no rows updated while updating cab status")
		}
	}
	return err
}

func (s *Services) UpdateCab(ctx context.Context, cab *models.Cab, tx *sqlx.Tx) error {
	res, err := tx.NamedExecContext(
		ctx,
		"UPDATE "+models.CabsTableName+ " "+
			"SET current_state=:current_state, last_updated_by=:last_updated_by, "+
			"cab_number=:cab_number, current_city_id=:current_city_id, "+
			"is_active=:is_active, last_ride_end_time=:last_ride_end_time "+
			"WHERE id=:id",
		cab,
	)
	if err == nil {
		totalRows, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if totalRows == 0 {
			return errors.New("no rows updated while updating cab status")
		}
	}
	return err
}
