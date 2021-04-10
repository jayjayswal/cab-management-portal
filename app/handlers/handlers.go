package handlers

import (
	controllers "cab-management-portal/app/controllers"
	"cab-management-portal/app/utilEntities"
	"log"
	"net/http"
)

type Handler struct {
	controller *controllers.Controller
	logger     *log.Logger
	//routeHelpers common.RouteHelpers
	//utilHelpers  common.UtilHelpers
}

func GetNewHandler(
	dependencies *utilEntities.Dependencies,
	controller *controllers.Controller,
) *Handler {
	return &Handler{
		logger: dependencies.Logger,
		controller: controller,
	}
}

func (h *Handler) WriteJSONResponse(w http.ResponseWriter, resp string, statusCode int) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, writeError := w.Write([]byte(resp))
	//_, writeError := w.Write([]byte(fmt.Sprintf("%v", resp)))
	if writeError != nil {
		h.logger.Print(writeError.Error())
	}
	return nil
}

func (h *Handler) WriteErrorResponse(w http.ResponseWriter, resp string, statusCode int) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, writeError := w.Write([]byte(resp))
	//_, writeError := w.Write([]byte(fmt.Sprintf("%v", resp)))
	if writeError != nil {
		h.logger.Print(writeError.Error())
	}
	return nil
}