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

func TestController_BookCab(t *testing.T) {
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
		payload *BookCabPayload
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   error
		want1  *models.Cab
		want2  *models.Ride
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
			got, got1, got2 := c.BookCab(tt.args.ctx, tt.args.payload)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BookCab() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BookCab() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("BookCab() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestController_CreateCabRequestEntry(t *testing.T) {
	type fields struct {
		appName   string
		tier      string
		dc        string
		services  servies.Services
		logger    *log.Logger
		validator *validator.Validate
	}
	type args struct {
		ctx    context.Context
		cityId int
		state  string
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
			if err := c.CreateCabRequestEntry(tt.args.ctx, tt.args.cityId, tt.args.state); (err != nil) != tt.wantErr {
				t.Errorf("CreateCabRequestEntry() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestController_FinishRide(t *testing.T) {
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
		payload *FinishRidePayload
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
			if err := c.FinishRide(tt.args.ctx, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("FinishRide() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestController_GetCityWiseRideInsight(t *testing.T) {
	type fields struct {
		appName   string
		tier      string
		dc        string
		services  servies.Services
		logger    *log.Logger
		validator *validator.Validate
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string][]RideInsight
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
			got, err := c.GetCityWiseRideInsight(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCityWiseRideInsight() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCityWiseRideInsight() got = %v, want %v", got, tt.want)
			}
		})
	}
}