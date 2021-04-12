package controllers

import (
	"cab-management-portal/app/models"
	"context"
	"errors"
	"time"
)

// This function books a cab
//Steps:
//Get city from DB
//If city is not active, reject cab booking
//Start DB Transaction:
//Get most idle cab in that city from DB
//If no cab found, Log request in CabRequest table with UNFULFILLED status
//If cab found, update cab status to ON_TRIP, and create ride entry with IN_PROGRESS status
//Log request in CabRequest table with FULFILLED status
//Complete the transaction
func (c *Controller) BookCab(ctx context.Context, payload *BookCabPayload) (*CabBooking, error) {
	err := c.validator.Struct(payload)
	if err != nil {
		return nil, err
	}
	city, err := c.services.GetCity(ctx, payload.CityId)
	if err != nil {
		return nil, err
	}
	if city.IsActive != 1 {
		return nil, errors.New("this city is not active for cab booking currently")
	}
	cab, ride, err := c.services.BookCabTxn(ctx, payload.CityId)
	if err != nil {
		return nil, err
	}
	if cab == nil {
		go c.CreateCabRequestEntry(context.Background(), payload.CityId, models.UnFulfilledRequestRideStatus)
		return nil, errors.New("no cabs found")
	} else {
		go c.CreateCabRequestEntry(context.Background(), payload.CityId, models.FulfilledRequestRideStatus)
		return &CabBooking{
			Message:      "Cab booked successfully",
			RideId:       ride.Id,
			CabId:        cab.Id,
			CabNo:        cab.CabNumber,
			StartCityId:  ride.StartCityId,
			StartTime:    ride.StartTime,
			EndTime:      ride.EndTime,
			CurrentState: ride.CurrentState,
		}, nil
	}
}

//This function finishes ride
//Steps:
//Start DB Transaction:
//Get Ride from DB
//If ride is not active, reject the request
//Get Cab from DB
//If cab is not active, reject the request
//Get most idle cab in that city from DB
//Update ride and cab with FINISHED and IDLE state
//Complete the transaction
func (c *Controller) FinishRide(ctx context.Context, payload *FinishRidePayload) (*CabBooking, error) {
	err := c.validator.Struct(payload)
	if err != nil {
		return nil, err
	}
	cab, ride, err := c.services.FinishRideTxn(ctx, payload.RideId)
	if err != nil {
		return nil, err
	}
	return &CabBooking{
		Message:      "Ride finished successfully",
		RideId:       ride.Id,
		CabId:        ride.CabId,
		CabNo:        cab.CabNumber,
		StartCityId:  ride.StartCityId,
		StartTime:    ride.StartTime,
		EndTime:      ride.EndTime,
		CurrentState: ride.CurrentState,
	}, nil
}

func (c *Controller) GetCityWiseRideInsight(ctx context.Context) (map[string][]RideInsight, error) {
	insights, err := c.services.GetCityWiseRideInsight(ctx)
	if err != nil {
		return nil, err
	}
	hoursMap := []string{"00:00AM to 04:00AM", "04:00AM to 08:00AM", "08:00AM to 12:00PM", "12:00PM to 04:00PM", "04:00PM to 08:00PM", "08:00PM to 00:00AM"}
	res := make(map[string][]RideInsight)
	for _, ins := range insights {
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

func (c *Controller) CreateCabRequestEntry(ctx context.Context, cityId int, state string) error {
	object := models.RideRequest{
		StartCityId:  cityId,
		CurrentState: state,
	}
	err := c.services.CreateRideRequest(ctx, &object)
	if err != nil {
		c.logger.Print(err.Error())
		return err
	}
	return nil
}

func (c *Controller) GetAllRides(ctx context.Context) ([]models.Ride, error) {
	return c.services.GetAllRides(ctx)
}

type RideInsight struct {
	Hours              string `json:"hours"`
	TotalRequests      int    `json:"total_requests"`
	FulfilledRequest   int    `json:"fulfilled_requests"`
	UnfulfilledRequest int    `json:"unfulfilled_requests"`
}

type CabBooking struct {
	Message      string     `json:"message"`
	RideId       int        `json:"ride_id"`
	CabId        int        `json:"cab_id"`
	CabNo        string     `json:"cab_number"`
	StartCityId  int        `json:"start_city_id"`
	StartTime    time.Time  `json:"start_time"`
	EndTime      *time.Time `json:"end_time"`
	CurrentState string     `json:"current_state"`
}

type BookCabPayload struct {
	CityId int `json:"city_id" validate:"required"`
}

type FinishRidePayload struct {
	RideId int `json:"ride_id" validate:"required"`
}
