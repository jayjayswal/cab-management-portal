package controllers

import (
	"cab-management-portal/app/models"
	"context"
	"errors"
)

// returns error or booked cab object
func (c *Controller) BookCab(ctx context.Context, payload *BookCabPayload) (error, *models.Cab, *models.Ride) {
	err := c.validator.Struct(payload)
	if err != nil {
		return err, nil, nil
	}
	city, err := c.services.GetCity(ctx, payload.CityId)
	if err != nil {
		return err, nil, nil
	}
	if city.IsActive != 1 {
		return errors.New("this city is not active for cab booking currently"), nil, nil
	}
	cab, ride, err := c.services.BookCabTxn(ctx, payload.CityId)
	if err != nil {
		return err, nil, nil
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
	err := c.validator.Struct(payload)
	if err != nil {
		return err
	}
	return c.services.FinishRide(ctx, payload.RideId)
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

type RideInsight struct {
	Hours              string `json:"hours"`
	TotalRequests      int    `json:"total_requests"`
	FulfilledRequest   int    `json:"fulfilled_requests"`
	UnfulfilledRequest int    `json:"unfulfilled_requests"`
}

type BookCabPayload struct {
	CityId int `json:"city_id" validate:"required"`
}

type FinishRidePayload struct {
	RideId int `json:"ride_id" validate:"required"`
}