package controllers

import (
	"cab-management-portal/app/models"
	"context"
)

func (c *Controller) CreateCity(ctx context.Context, payload *CreateCityPayload) error {
	err := c.validator.Struct(payload)
	if err != nil {
		return err
	}
	city := models.City{
		Name:          payload.Name,
		IsActive:      1,
	}
	return c.services.CreateCity(ctx, &city)
}

func (c *Controller) GetAllCities(ctx context.Context) ([]models.City, error) {
	return c.services.GetAllCities(ctx)
}

type CreateCityPayload struct {
	Name string `json:"name" validate:"required,min=2,max=50"`
}
