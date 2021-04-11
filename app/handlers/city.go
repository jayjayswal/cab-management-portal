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
	_ = h.WriteJSONResponse(writer, `{"message":"created"}`, http.StatusOK)
}
