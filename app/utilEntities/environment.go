package utilEntities

import (
	"cab-management-portal/app/constants"
	"log"
	"os"
	"strconv"
)

type Environment struct {
	AppName        string
	Tier           string
	Port           int
}


func GetEnvironment(l *log.Logger) *Environment {
	var port string
	var tier string
	var appName string


	appName, isNamePresent := os.LookupEnv(constants.AppName)
	if !isNamePresent {
		l.Panic("App name not found in Environment")
	}
	port, isPortPresent := os.LookupEnv(constants.Port)
	if !isPortPresent {
		port = constants.DefaultPort
	}
	tier, isTierPresent := os.LookupEnv(constants.Tier)
	if !isTierPresent {
		tier = constants.DefaultTier
	}
	portInt, err := strconv.Atoi(port)
	if err != nil {
		l.Panic("Port in non integer, please give correct port")
	}

	return &Environment{
		AppName: appName,
		Tier:    tier,
		Port:    portInt,
	}
}