package handlers

import (
	"cab-management-portal/app/controllers"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (h *Handler) CreateCity(writer http.ResponseWriter, request *http.Request) {
	reqBody, _ := ioutil.ReadAll(request.Body)
	var obj controllers.CreateCityPayload
	err := json.Unmarshal(reqBody, &obj)
	if err != nil {
		h.logger.Print(err.Error())
		_ = h.Write500ErrorResponse(writer, err)
		return
	}
	err = h.controller.CreateCity(context.Background(), &obj)
	if err != nil {
		h.logger.Print(err.Error())
		_ = h.Write500ErrorResponse(writer, err)
		return
	}
	_ = h.WriteJSONResponse(writer, `{"message":"city created"}`, http.StatusOK)
}

func (h *Handler) GetAllCities(writer http.ResponseWriter, request *http.Request) {
	cities, err := h.controller.GetAllCities(context.Background())
	if err != nil {
		h.logger.Print(err.Error())
		_ = h.Write500ErrorResponse(writer, err)
		return
	}
	res, err := json.Marshal(cities)
	if err != nil {
		h.logger.Print(err.Error())
		_ = h.Write500ErrorResponse(writer, err)
		return
	}
	_ = h.WriteJSONResponse(writer, string(res), http.StatusOK)
}