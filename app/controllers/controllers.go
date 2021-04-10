package controllers

import (
	"cab-management-portal/app/servies"
	"cab-management-portal/app/utilEntities"
	"log"
)

type Controller struct {
	appName  string
	tier     string
	dc       string
	services *servies.Services
	logger   *log.Logger
	//helpers  common.UtilHelpers
}

func GetNewController(
	dependencies *utilEntities.Dependencies,
	s *servies.Services,
) *Controller {
	return &Controller{
		services: s,
		appName: dependencies.AppName,
		tier:    dependencies.Tier,
		logger:  dependencies.Logger,
		//helpers:  jobsApp.UtilHelpers,
	}
}
