package controllers

import (
	"cab-management-portal/app/models"
	"context"
)

func (c *Controller) CreateCab(ctx context.Context, payload *CreateCabPayload) error {
	cab := models.Cab{
		CabNumber:     payload.CabNumber,
		CurrentState:  models.CabIdleState,
		CurrentCityId: payload.CurrentCityId,
		IsActive:      1,
		LastUpdatedBy: 1,
	}
	return c.services.CreateCab(ctx, &cab)
}

func (c *Controller) UpdateCabCity(ctx context.Context, payload *UpdateCityPayload) error {
	cab := models.Cab{
		Id: payload.CabId,
		CurrentCityId: &payload.CurrentCityId,
	}
	return c.services.UpdateCabCity(ctx, &cab)
}

type CreateCabPayload struct {
	CabNumber     string `json:"cab_number"`
	CurrentCityId *int   `json:"current_city_id"`
}

type UpdateCityPayload struct {
	CabId         int `json:"cab_id"`
	CurrentCityId int `json:"current_city_id"`
}
