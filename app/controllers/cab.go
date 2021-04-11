package controllers

import (
	"cab-management-portal/app/models"
	"context"
	"errors"
)

func (c *Controller) CreateCab(ctx context.Context, payload *CreateCabPayload) error {
	cab := models.Cab{
		CabNumber:     payload.CabNumber,
		CurrentState:  models.CabIdleState,
		CurrentCityId: payload.CurrentCityId,
		IsActive:      1,
	}
	return c.services.CreateCab(ctx, &cab)
}

func (c *Controller) UpdateCabCity(ctx context.Context, payload *UpdateCityPayload) error {
	tx := c.services.Sequel.MustBegin()
	cab, err := c.services.GetCabForUpdate(ctx, payload.CabId, tx)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	if cab.CurrentState == models.CabOnTripState {
		_ = tx.Rollback()
		return errors.New("cab is on trip state, finish the ride before you change the city")
	}
	cab.CurrentCityId = &payload.CurrentCityId
	err = c.services.UpdateCabCity(ctx, cab, tx)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	_ = tx.Commit()
	return nil
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
