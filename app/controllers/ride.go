package controllers

import (
	"cab-management-portal/app/models"
	"context"
	"errors"
	"time"
)

// returns error or booked cab object
func (c *Controller) BookCab(ctx context.Context, payload *BookCabPayload) (error, *models.Cab, *models.Ride) {
	city, err := c.services.GetCity(ctx, payload.CityId)
	if err != nil {
		return err, nil, nil
	}
	if city.IsActive != 1 {
		return errors.New("this city is not active for cab booking currently"), nil, nil
	}
	tx := c.services.Sequel.MustBegin()
	var cab *models.Cab = nil
	var ride *models.Ride = nil
	cabs, err := c.services.GetMostIdleCabOfCity(ctx, city.Id, tx)
	if err != nil {
		_ = tx.Rollback()
		return err, nil, nil
	}
	if cabs!= nil && len(cabs) >= 0 {
		cab = &cabs[0]
		ride = &models.Ride{
			CabId:         cab.Id,
			StartCityId:   payload.CityId,
			CurrentState:  models.InProgressRideStatus,
			LastUpdatedBy: 1,
		}
		err = c.services.CreateRide(ctx, ride, tx)
		if err != nil {
			_ = tx.Rollback()
			return err, nil, nil
		}
		cab.CurrentState = models.CabOnTripState
		cab.LastUpdatedBy = 1
		cab.CurrentCityId = nil
		err = c.services.UpdateCabState(ctx, cab, tx)
		if err != nil {
			_ = tx.Rollback()
			return err, nil, nil
		}
		_ = tx.Commit()
	} else {
		_ = tx.Rollback()
	}
	if cab == nil {
		go c.CreateCabRequestEntry(context.Background(), payload.CityId, models.UnFulfilledRequestRideStatus)
		return errors.New("no cabs found"), nil, nil
	} else {
		go c.CreateCabRequestEntry(context.Background(), payload.CityId, models.FulfilledRequestRideStatus)
		return nil, cab, ride
	}
}

func (c *Controller) FinishRide(ctx context.Context, payload *FinishRidePayload) error {
	tx := c.services.Sequel.MustBegin()
	now := time.Now()
	ride, err := c.services.GetRideForUpdate(ctx, payload.RideId, tx)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	if ride.CurrentState != models.InProgressRideStatus {
		_ = tx.Rollback()
		return errors.New("this ride is not in progress anymore")
	}
	cab, err := c.services.GetCabForUpdate(ctx, ride.CabId, tx)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	if cab.CurrentState != models.CabOnTripState {
		_ = tx.Rollback()
		return errors.New("cab is not in trip state")
	}
	ride.CurrentState = models.FinishedRideStatus
	ride.LastUpdatedBy = 1
	ride.EndTime = &now
	err = c.services.UpdateRide(ctx, ride, tx)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	cab.LastUpdatedBy = 1
	cab.CurrentState = models.CabIdleState
	cab.LastRideEndTime = &now
	err = c.services.UpdateCab(ctx, cab, tx)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	_ = tx.Commit()
	return nil
}

func (c *Controller) CreateCabRequestEntry(ctx context.Context, cityId int, state string) {
	object := models.RideRequest{
		StartCityId:  cityId,
		CurrentState: state,
	}
	err := c.services.CreateRideRequest(ctx, &object)
	if err != nil {
		c.logger.Print(err.Error())
	}
}

type BookCabPayload struct {
	CityId int `json:"city_id"`
}

type FinishRidePayload struct {
	RideId int `json:"ride_id"`
}