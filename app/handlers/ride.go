package handlers

import (
	"cab-management-portal/app/controllers"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (h *Handler) RequestRide(writer http.ResponseWriter, request *http.Request) {
	reqBody, _ := ioutil.ReadAll(request.Body)
	var obj controllers.BookCabPayload
	err := json.Unmarshal(reqBody, &obj)
	if err != nil {
		h.logger.Print(err.Error())
		_ = h.Write500ErrorResponse(writer, err)
		return
	}
	err, _, _ = h.controller.BookCab(context.Background(), &obj)
	if err != nil {
		h.logger.Print(err.Error())
		_ = h.Write500ErrorResponse(writer, err)
		return
	}
	_ = h.WriteJSONResponse(writer, `{"message":"Cab Booked"}`, http.StatusOK)
}

func (h *Handler) FinishRide(writer http.ResponseWriter, request *http.Request) {
	reqBody, _ := ioutil.ReadAll(request.Body)
	var obj controllers.FinishRidePayload
	err := json.Unmarshal(reqBody, &obj)
	if err != nil {
		h.logger.Print(err.Error())
		_ = h.Write500ErrorResponse(writer, err)
		return
	}
	err = h.controller.FinishRide(context.Background(), &obj)
	if err != nil {
		h.logger.Print(err.Error())
		_ = h.Write500ErrorResponse(writer, err)
		return
	}
	_ = h.WriteJSONResponse(writer, `{"message":"Ride Finished"}`, http.StatusOK)
}
func (h *Handler) GetCityWiseRideInsight(writer http.ResponseWriter, request *http.Request) {
	insights, err := h.controller.GetCityWiseRideInsight(context.Background())
	if err != nil {
		h.logger.Print(err.Error())
		_ = h.Write500ErrorResponse(writer, err)
		return
	}
	res, err := json.Marshal(insights)
	if err != nil {
		h.logger.Print(err.Error())
		_ = h.Write500ErrorResponse(writer, err)
		return
	}
	_ = h.WriteJSONResponse(writer, string(res), http.StatusOK)
}
