package handlers

import (
	controllers "cab-management-portal/app/controllers"
	"cab-management-portal/app/utilEntities"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
)

type Handler struct {
	controller *controllers.Controller
	logger     *log.Logger
	validator  *validator.Validate
	//routeHelpers common.RouteHelpers
	//utilHelpers  common.UtilHelpers
}

func GetNewHandler(
	dependencies *utilEntities.Dependencies,
	controller *controllers.Controller,
) *Handler {
	return &Handler{
		logger:     dependencies.Logger,
		controller: controller,
		validator:  dependencies.Validator,
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

func (h *Handler) Write500ErrorResponse(w http.ResponseWriter, err error) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(500)
	_, writeError := w.Write([]byte(`{
		"message":"Something went wrong, Please try again",
		"error_message": "` + err.Error() + `"
	}`))
	if writeError != nil {
		h.logger.Print(writeError.Error())
	}
	return nil
}

func (h *Handler) Write404ErrorResponse(w http.ResponseWriter) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(404)
	_, writeError := w.Write([]byte(`{"message":"Object Not found, Try again later"}`))
	if writeError != nil {
		h.logger.Print(writeError.Error())
	}
	return nil
}
