package utilEntities

import (
	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
	"log"
)

type Dependencies struct {
	AppName      string
	Logger       *log.Logger
	ConfigHelper *ConfigHelper
	Router       *mux.Router
	Validator    *validator.Validate
	//RouteHelpers RouteHelpers
	//UtilHelpers  UtilHelpers
	Port int
	Tier string
}

func GetDependencies(environment *Environment, logger *log.Logger) *Dependencies {
	r := mux.NewRouter()
	confHelper := GetConfigHelper(environment)
	validate := validator.New()
	return &Dependencies{
		AppName:      environment.AppName,
		Logger:       logger,
		Validator:    validate,
		ConfigHelper: confHelper,
		Port:         environment.Port,
		Tier:         environment.Tier,
		Router:       r,
	}
}
