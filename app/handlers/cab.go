package handlers

import (
	"cab-management-portal/app/controllers"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (h *Handler) CreateCab(writer http.ResponseWriter, request *http.Request) {
	reqBody, _ := ioutil.ReadAll(request.Body)
	var obj controllers.CreateCabPayload
	err := json.Unmarshal(reqBody, &obj)
	if err != nil {
		h.logger.Print(err.Error())
		_ = h.Write500ErrorResponse(writer, err)
		return
	}
	err = h.controller.CreateCab(context.Background(), &obj)
	if err != nil {
		h.logger.Print(err.Error())
		_ = h.Write500ErrorResponse(writer, err)
		return
	}
	_ = h.WriteJSONResponse(writer, `{"message":"cab created"}`, http.StatusOK)
}

func (h *Handler) GetAllCabs(writer http.ResponseWriter, request *http.Request) {
	cabs, err := h.controller.GetAllCabs(context.Background())
	if err != nil {
		h.logger.Print(err.Error())
		_ = h.Write500ErrorResponse(writer, err)
		return
	}
	res, err := json.Marshal(cabs)
	if err != nil {
		h.logger.Print(err.Error())
		_ = h.Write500ErrorResponse(writer, err)
		return
	}
	_ = h.WriteJSONResponse(writer, string(res), http.StatusOK)
}

func (h *Handler) UpdateCity(writer http.ResponseWriter, request *http.Request) {
	reqBody, _ := ioutil.ReadAll(request.Body)
	var obj controllers.UpdateCityPayload
	err := json.Unmarshal(reqBody, &obj)
	if err != nil {
		h.logger.Print(err.Error())
		_ = h.Write500ErrorResponse(writer, err)
		return
	}
	err = h.controller.UpdateCabCity(context.Background(), &obj)
	if err != nil {
		h.logger.Print(err.Error())
		_ = h.Write500ErrorResponse(writer, err)
		return
	}
	_ = h.WriteJSONResponse(writer, `{"message":"cab updated"}`, http.StatusOK)
}

func (h *Handler) GetCabActivities(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	var id = vars["cabId"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.logger.Print(err.Error())
		_ = h.Write500ErrorResponse(writer, err)
		return
	}
	activities, err := h.controller.GetCabActivities(context.Background(), idInt)
	if err != nil {
		h.logger.Print(err.Error())
		_ = h.Write500ErrorResponse(writer, err)
		return
	}
	res, err := json.Marshal(activities)
	if err != nil {
		h.logger.Print(err.Error())
		_ = h.Write500ErrorResponse(writer, err)
		return
	}
	_ = h.WriteJSONResponse(writer, string(res), http.StatusOK)
}

func (h *Handler) GetCabIdleTime(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	var id = vars["cabId"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.logger.Print(err.Error())
		_ = h.Write500ErrorResponse(writer, err)
		return
	}
	idleHours, err := h.controller.GetCabIdleTime(context.Background(), idInt)
	if err != nil {
		h.logger.Print(err.Error())
		_ = h.Write500ErrorResponse(writer, err)
		return
	}
	res, err := json.Marshal(idleHours)
	if err != nil {
		h.logger.Print(err.Error())
		_ = h.Write500ErrorResponse(writer, err)
		return
	}
	_ = h.WriteJSONResponse(writer, string(res), http.StatusOK)
}