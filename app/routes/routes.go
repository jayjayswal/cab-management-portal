package routes

import (
	controllers "cab-management-portal/app/controllers"
	"cab-management-portal/app/handlers"
	"cab-management-portal/app/servies"
	"cab-management-portal/app/utilEntities"
	"github.com/gorilla/mux"
	"net/http"
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

	dependencies.Router.HandleFunc("/health-check",handler.HealthCheck)

	apiRouter := dependencies.Router.PathPrefix("/api").Subrouter()
	amw := utilEntities.AuthenticationMiddleware{}
	amw.Populate()
	apiRouter.Use(amw.Middleware)

	cityRouter := apiRouter.PathPrefix("/city").Subrouter()
	cityRouter.HandleFunc("/create", handler.CreateCity).Methods(http.MethodPost)
	//articlesRouter.HandleFunc("/all", handler.GetAllCities).Methods(http.MethodGet)
	//articlesRouter.HandleFunc("/{articlesId}", handler.GetCity).Methods(http.MethodGet)
	//articlesRouter.HandleFunc("/update", handler.UpdateCity).Methods(http.MethodPut)
	//articlesRouter.HandleFunc("/{articlesId}", handler.DeleteCity).Methods(http.MethodDelete)

	cabRouter := apiRouter.PathPrefix("/cab").Subrouter()
	cabRouter.HandleFunc("/create", handler.CreateCab).Methods(http.MethodPost)
	cabRouter.HandleFunc("/update-city", handler.UpdateCity).Methods(http.MethodPut)
	cabRouter.HandleFunc("/recent-activities/{cabId}", handler.GetCabActivities).Methods(http.MethodGet)
	//articlesRouter.HandleFunc("/all", handler.GetAllCities).Methods(http.MethodGet)
	//articlesRouter.HandleFunc("/{articlesId}", handler.GetCity).Methods(http.MethodGet)
	//articlesRouter.HandleFunc("/update", handler.UpdateCity).Methods(http.MethodPut)
	//articlesRouter.HandleFunc("/{articlesId}", handler.DeleteCity).Methods(http.MethodDelete)

	rideRouter := apiRouter.PathPrefix("/ride").Subrouter()
	rideRouter.HandleFunc("/request-new-ride", handler.RequestRide).Methods(http.MethodPost)
	rideRouter.HandleFunc("/finish-ride", handler.FinishRide).Methods(http.MethodPost)
	rideRouter.HandleFunc("/insights", handler.GetCityWiseRideInsight).Methods(http.MethodGet)

	return dependencies.Router
}