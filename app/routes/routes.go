package routes

import (
	controllers "cab-management-portal/app/controllers"
	"cab-management-portal/app/handlers"
	"cab-management-portal/app/servies"
	"cab-management-portal/app/utilEntities"
	"github.com/gorilla/mux"
)

func Init(dependencies *utilEntities.Dependencies) *mux.Router {

	services, err := servies.GetServiceObject(dependencies)
	if err != nil {
		dependencies.Logger.Panic(err.Error())
	}
	controller := controllers.GetNewController(
		dependencies,
		services,
	)
	handler := handlers.GetNewHandler(
		dependencies,
		controller,
	)

	dependencies.Router.HandleFunc("/",handler.HealthCheck)
	//_ := dependencies.Router.PathPrefix("/api").Subrouter()


	return dependencies.Router
}