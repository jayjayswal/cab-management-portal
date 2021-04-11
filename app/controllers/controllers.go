package controllers

import (
	"cab-management-portal/app/servies"
	"cab-management-portal/app/utilEntities"
	"gopkg.in/go-playground/validator.v9"
	"log"
)

type Controller struct {
	appName  string
	tier     string
	dc       string
	services servies.Services
	logger   *log.Logger
	validator  *validator.Validate
	//helpers  common.UtilHelpers
}

func GetNewController(
	dependencies *utilEntities.Dependencies,
	s servies.Services,
) *Controller {
	return &Controller{
		services: s,
		appName: dependencies.AppName,
		tier:    dependencies.Tier,
		logger:  dependencies.Logger,
		validator: dependencies.Validator,
		//helpers:  jobsApp.UtilHelpers,
	}
}
