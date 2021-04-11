package servies

import (
	"cab-management-portal/app/models"
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

func (s *Services) GetRide(ctx context.Context, id int) (*models.Ride, error) {
	ride := models.Ride{}
	query := "SELECT * FROM " +
		models.RidesTableName +
		" WHERE id=?"
	err := s.Sequel.GetContext(ctx, &ride, query, id)
	if err != nil {
		return nil, err
	}
	return &ride, nil
}

func (s *Services) GetRideForUpdate(ctx context.Context, id int, tx *sqlx.Tx) (*models.Ride, error) {
	ride := models.Ride{}
	query := "SELECT * FROM " +
		models.RidesTableName +
		" WHERE id=? FOR UPDATE"
	err := tx.GetContext(ctx, &ride, query, id)
	if err != nil {
		return nil, err
	}
	return &ride, nil
}

func (s *Services) CreateRide(ctx context.Context, ride *models.Ride, tx *sqlx.Tx) error {
	query := `INSERT INTO ` + models.RidesTableName +
		` (cab_id,start_city_id,current_state) VALUES ` +
		`(:cab_id,:start_city_id,:current_state)`
	var err error
	var res sql.Result
	if tx != nil {
		res, err = tx.NamedExecContext(ctx, query, ride)
	} else {
		res, err = s.Sequel.NamedExecContext(ctx, query, ride)
	}
	if err == nil {
		id, err := res.LastInsertId()
		if err != nil {
			ride.Id = int(id)
		} else {
			return err
		}
	}
	return err
}

func (s *Services) UpdateRide(ctx context.Context, ride *models.Ride, tx *sqlx.Tx) error {
	res, err := tx.NamedExecContext(
		ctx,
		"UPDATE "+models.RidesTableName+" "+
			"SET cab_id=:cab_id, start_city_id=:start_city_id, "+
			"current_state=:current_state, "+
			"start_time=:start_time, end_time=:end_time "+
			"WHERE id=:id",
		ride,
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

func (s *Services) CreateRideRequest(ctx context.Context, rideRequest *models.RideRequest) error {
	res, err := s.Sequel.NamedExecContext(
		ctx,
		`INSERT INTO `+models.RideRequestTableName+
			` (start_city_id,current_state) VALUES `+
			`(:start_city_id,:current_state)`,
		rideRequest,
	)
	if err == nil {
		id, err := res.LastInsertId()
		if err != nil {
			rideRequest.Id = int(id)
		} else {
			return err
		}
	}
	return err
}

func (s *Services) GetCityWiseRideInsight(ctx context.Context) ([]RideInsight, error) {
	var rideInsights []RideInsight
	query := "SELECT c.name, rr.start_city_id, floor(hour(rr.created) / 4) as hourGroupId, count(1) as total_requests, " +
		"SUM(IF(rr.current_state=\"FULFILLED\", 1, 0)) as fulfilled_requests, " +
		"SUM(IF(rr.current_state=\"UNFULFILLED\", 1, 0)) as unfulfilled_requests " +
		"FROM ride_requests as rr LEFT JOIN cities as c ON c.id = rr.start_city_id " +
		"WHERE rr.created > DATE_SUB(now(), INTERVAL 30 DAY) " +
		"GROUP BY rr.start_city_id, hourgroupId " +
		"ORDER BY total_requests DESC " +
		"LIMIT 100 "
	err := s.Sequel.SelectContext(ctx, &rideInsights, query)
	if err != nil {
		return nil, err
	}
	if rideInsights == nil {
		rideInsights = []RideInsight{}
	}
	return rideInsights, nil
}

type RideInsight struct {
	CityName           string `db:"name"`
	StartCityId        int    `db:"start_city_id"`
	HourGroupId        int    `db:"hourGroupId"`
	TotalRequests      int    `db:"total_requests"`
	FulfilledRequest   int    `db:"fulfilled_requests"`
	UnfulfilledRequest int    `db:"unfulfilled_requests"`
}

/**
SELECT c.name, rr.start_city_id,
floor(hour(rr.created) / 4) as hourGroupId,
count(1) as total_requests,
SUM(IF(rr.current_state="FULFILLED", 1, 0)) as fulfilled_requests, SUM(IF(rr.current_state="UNFULFILLED", 1, 0)) as unfulfilled_requests
FROM ride_requests as rr
LEFT JOIN cities as c ON c.id = rr.start_city_id
WHERE rr.created > DATE_SUB(now(), INTERVAL 30 DAY)
GROUP BY rr.start_city_id, hourGroupId
ORDER BY total_requests DESC
LIMIT 100
*/
