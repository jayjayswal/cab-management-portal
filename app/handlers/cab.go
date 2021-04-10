package handlers

import (
	"cab-management-portal/app/controllers"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (h *Handler) CreateCab(writer http.ResponseWriter, request *http.Request) {
	reqBody, _ := ioutil.ReadAll(request.Body)
	var obj controllers.CreateCabPayload
	err := json.Unmarshal(reqBody, &obj)
	if err != nil {
		h.logger.Print(err.Error())
		_ = h.Write500ErrorResponse(writer)
		return
	}
	err = h.controller.CreateCab(context.Background(),&obj)
	if err != nil {
		h.logger.Print(err.Error())
		_ = h.Write500ErrorResponse(writer)
		return
	}
	_ = h.WriteJSONResponse(writer, `{"message":"created"}`, http.StatusOK)
}