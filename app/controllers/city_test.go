package controllers

import (
	"cab-management-portal/app/servies"
	mockService "cab-management-portal/app/servies/mock"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"os"
	"testing"
)

func TestController_CreateCity(t *testing.T) {
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
		payload *CreateCityPayload
	}
	ctrl := gomock.NewController(t)
	validate := validator.New()
	logger := log.New(os.Stderr, "", log.LstdFlags)

	mockServiceErr := mockService.NewMockServices(ctrl)
	mockServiceErr.EXPECT().CreateCity(gomock.Any(), gomock.Any()).Return(errors.New("error from DB"))
	mockService := mockService.NewMockServices(ctrl)
	mockService.EXPECT().CreateCity(gomock.Any(), gomock.Any()).Return(nil)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Valid Name",
			fields: fields{
				services:  mockService,
				logger:    logger,
				validator: validate,
			},
			args: args{
				ctx: context.Background(),
				payload: &CreateCityPayload{
					Name: "test",
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid Name",
			fields: fields{
				services:  mockService,
				logger:    logger,
				validator: validate,
			},
			args: args{
				ctx: context.Background(),
				payload: &CreateCityPayload{
					Name: "t",
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid Name",
			fields: fields{
				services:  mockServiceErr,
				logger:    logger,
				validator: validate,
			},
			args: args{
				ctx: context.Background(),
				payload: &CreateCityPayload{
					Name: "test",
				},
			},
			wantErr: true,
		},
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
			if err := c.CreateCity(tt.args.ctx, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("CreateCity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
