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

func (c *Controller) GetCityWiseRideInsight(ctx context.Context) (map[string][]RideInsight, error) {
	insights, err := c.services.GetCityWiseRideInsight(ctx)
	if err != nil {
		return nil, err
	}
	hoursMap := []string{"00:00AM to 04:00AM","04:00AM to 08:00AM","08:00AM to 12:00PM","12:00PM to 04:00PM","04:00PM to 08:00PM","08:00PM to 00:00AM"}
	res := make(map[string][]RideInsight)
	for _, ins :=  range insights {
		insArr, ok := res[ins.CityName]
		if !ok {
			insArr = []RideInsight{}
			res[ins.CityName] = insArr
		}
		hour := "UNKNOWN"
		if len(hoursMap) >= ins.HourGroupId {
			hour = hoursMap[ins.HourGroupId]
		}
		res[ins.CityName] = append(insArr, RideInsight{
			Hours:              hour,
			TotalRequests:      ins.TotalRequests,
			FulfilledRequest:   ins.FulfilledRequest,
			UnfulfilledRequest: ins.UnfulfilledRequest,
		})
	}
	return res, nil
}

type RideInsight struct {
	Hours              string `json:"hours"`
	TotalRequests      int    `json:"total_requests"`
	FulfilledRequest   int    `json:"fulfilled_requests"`
	UnfulfilledRequest int    `json:"unfulfilled_requests"`
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