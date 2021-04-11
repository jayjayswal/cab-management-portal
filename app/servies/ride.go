package servies

import (
	"cab-management-portal/app/models"
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"time"
)

func (s *Service) GetRide(ctx context.Context, id int) (*models.Ride, error) {
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

func (s *Service) GetAllRides(ctx context.Context) ([]models.Ride, error) {
	var rides []models.Ride
	err := s.Sequel.SelectContext(ctx, &rides, "SELECT * FROM "+
		models.RidesTableName)
	if err != nil {
		return nil, err
	}
	if rides == nil {
		rides = []models.Ride{}
	}
	return rides, nil
}

func (s *Service) GetRideForUpdate(ctx context.Context, id int, tx *sqlx.Tx) (*models.Ride, error) {
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

func (s *Service) CreateRide(ctx context.Context, ride *models.Ride, tx *sqlx.Tx) error {
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
			return err
		}
		ride.Id = int(id)
	}
	return err
}

func (s *Service) UpdateRide(ctx context.Context, ride *models.Ride, tx *sqlx.Tx) error {
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

func (s *Service) CreateRideRequest(ctx context.Context, rideRequest *models.RideRequest) error {
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

func (s *Service) GetCityWiseRideInsight(ctx context.Context) ([]RideInsight, error) {
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

func (s *Service) BookCabTxn(ctx context.Context, cityId int) (*models.Cab, *models.Ride, error) {
	tx := s.Sequel.MustBegin()
	var cab *models.Cab = nil
	var ride *models.Ride = nil
	cabs, err := s.GetMostIdleCabOfCity(ctx, cityId, tx)
	if err != nil {
		_ = tx.Rollback()
		return nil, nil, err
	}
	if cabs != nil && len(cabs) >= 0 {
		cab = &cabs[0]
		ride = &models.Ride{
			CabId:        cab.Id,
			StartCityId:  cityId,
			CurrentState: models.InProgressRideStatus,
		}
		err = s.CreateRide(ctx, ride, tx)
		if err != nil {
			_ = tx.Rollback()
			return nil, nil, err
		}
		cab.CurrentState = models.CabOnTripState
		cab.CurrentCityId = nil
		err = s.UpdateCabState(ctx, cab, tx)
		if err != nil {
			_ = tx.Rollback()
			return nil, nil, err
		}
		_ = tx.Commit()
	} else {
		_ = tx.Rollback()
	}
	return cab, ride, nil
}

func (s *Service) FinishRideTxn(ctx context.Context, rideId int) (*models.Cab, *models.Ride, error) {
	tx := s.Sequel.MustBegin()
	now := time.Now()
	ride, err := s.GetRideForUpdate(ctx, rideId, tx)
	if err != nil {
		_ = tx.Rollback()
		return nil, nil, err
	}
	if ride.CurrentState != models.InProgressRideStatus {
		_ = tx.Rollback()
		return nil, nil, errors.New("this ride is not in progress anymore")
	}
	cab, err := s.GetCabForUpdate(ctx, ride.CabId, tx)
	if err != nil {
		_ = tx.Rollback()
		return nil, nil, err
	}
	if cab.CurrentState != models.CabOnTripState {
		_ = tx.Rollback()
		return nil, nil, errors.New("cab is not in trip state")
	}
	ride.CurrentState = models.FinishedRideStatus
	ride.EndTime = &now
	err = s.UpdateRide(ctx, ride, tx)
	if err != nil {
		_ = tx.Rollback()
		return nil, nil, err
	}
	cab.CurrentState = models.CabIdleState
	cab.LastRideEndTime = &now
	err = s.UpdateCab(ctx, cab, tx)
	if err != nil {
		_ = tx.Rollback()
		return nil, nil, err
	}
	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return nil, nil, err
	}
	return cab, ride, nil
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
