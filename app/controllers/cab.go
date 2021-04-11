package controllers

import (
	"cab-management-portal/app/models"
	"context"
	"errors"
)

func (c *Controller) CreateCab(ctx context.Context, payload *CreateCabPayload) error {
	err := c.validator.Struct(payload)
	if err != nil {
		return err
	}
	cab := models.Cab{
		CabNumber:     payload.CabNumber,
		CurrentState:  models.CabIdleState,
		CurrentCityId: payload.CurrentCityId,
		IsActive:      1,
	}
	return c.services.CreateCab(ctx, &cab)
}

func (c *Controller) UpdateCabCity(ctx context.Context, payload *UpdateCityPayload) error {
	err := c.validator.Struct(payload)
	if err != nil {
		return err
	}
	cab, err := c.services.GetCab(ctx, payload.CabId)
	if err != nil {
		return err
	}
	if cab.CurrentState == models.CabOnTripState {
		return errors.New("cab is on trip state, finish the ride before you change the city")
	}
	return c.services.UpdateCabCityTxn(ctx, payload.CabId, payload.CurrentCityId)
}

func (c *Controller) GetCabActivities(ctx context.Context, cabId int) ([]models.CabAudit, error) {
	return c.services.GetCabActivities(ctx, cabId)
}

type CreateCabPayload struct {
	CabNumber     string `json:"cab_number" validate:"required,min=6,max=15"`
	CurrentCityId *int   `json:"current_city_id"`
}

type UpdateCityPayload struct {
	CabId         int `json:"cab_id" validate:"required"`
	CurrentCityId int `json:"current_city_id" validate:"required"`
}
