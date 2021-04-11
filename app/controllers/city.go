package controllers

import (
	"cab-management-portal/app/models"
	"context"
)

func (c *Controller) CreateCity(ctx context.Context, payload *CreateCityPayload) error {
	city := models.City{
		Name:          payload.Name,
		IsActive:      1,
	}
	return c.services.CreateCity(ctx, &city)
}

type CreateCityPayload struct {
	Name string `json:"name" validate:"required,min=2,max=50"`
}
