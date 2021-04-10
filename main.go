package main

import (
	"cab-management-portal/app/routes"
	"cab-management-portal/app/utilEntities"
	"log"
	"net/http"
	"os"
	"strconv"
)



func main() {
	logger := log.New(os.Stderr, "", log.LstdFlags)
	environment := utilEntities.GetEnvironment(logger)
	dependencies := utilEntities.GetDependencies(environment, logger)
 	routes.Init(dependencies)
	startServer(dependencies)
}

func startServer(dependencies *utilEntities.Dependencies)  {
	p := strconv.Itoa(dependencies.Port)
	dependencies.Logger.Print("Listening at " + p)
	err := http.ListenAndServe(":"+p, dependencies.Router)
	if err != nil {
		dependencies.Logger.Print("Error While starting server: " + err.Error())
		os.Exit(1)
	}
}