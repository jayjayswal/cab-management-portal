package controllers

import (
	"cab-management-portal/app/models"
	"cab-management-portal/app/servies"
	"context"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"reflect"
	"testing"
)

func TestController_CreateCab(t *testing.T) {
	type fields struct {
		appName   string
		tier      string
		dc        string
		services  servies.Services
		logger    *log.Logger
		validator *validator.Validate
	}
	type args struct {
		ctx     context.Context
		payload *CreateCabPayload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Controller{
				appName:   tt.fields.appName,
				tier:      tt.fields.tier,
				dc:        tt.fields.dc,
				services:  tt.fields.services,
				logger:    tt.fields.logger,
				validator: tt.fields.validator,
			}
			if err := c.CreateCab(tt.args.ctx, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("CreateCab() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestController_GetCabActivities(t *testing.T) {
	type fields struct {
		appName   string
		tier      string
		dc        string
		services  servies.Services
		logger    *log.Logger
		validator *validator.Validate
	}
	type args struct {
		ctx   context.Context
		cabId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.CabAudit
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Controller{
				appName:   tt.fields.appName,
				tier:      tt.fields.tier,
				dc:        tt.fields.dc,
				services:  tt.fields.services,
				logger:    tt.fields.logger,
				validator: tt.fields.validator,
			}
			got, err := c.GetCabActivities(tt.args.ctx, tt.args.cabId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCabActivities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCabActivities() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestController_UpdateCabCity(t *testing.T) {
	type fields struct {
		appName   string
		tier      string
		dc        string
		services  servies.Services
		logger    *log.Logger
		validator *validator.Validate
	}
	type args struct {
		ctx     context.Context
		payload *UpdateCityPayload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Controller{
				appName:   tt.fields.appName,
				tier:      tt.fields.tier,
				dc:        tt.fields.dc,
				services:  tt.fields.services,
				logger:    tt.fields.logger,
				validator: tt.fields.validator,
			}
			if err := c.UpdateCabCity(tt.args.ctx, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("UpdateCabCity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
